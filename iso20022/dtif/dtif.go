package dtif

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"math/big"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/pkg/errors"

	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
)

type Dtif struct {
	log                  logger.Logger
	distributedLedger    string
	sourceAddress        string
	username             string
	password             string
	dtiToDenom           map[string]string
	dtiToPriceMultiplier map[string]*big.Int
	denomToDti           map[string]string
	lastVersion          string
	lock                 sync.RWMutex
}

type DigitalToken interface {
	DTI() string
	Denom() *string
	PriceMultiplier() *big.Int
}

type DigitalTokenType [2]int

var (
	AuxiliaryDigitalToken                                 DigitalTokenType = [2]int{0, 0}
	NativeDigitalTokenBlockchain                          DigitalTokenType = [2]int{1, 1}
	NativeDigitalTokenOther                               DigitalTokenType = [2]int{1, 0}
	DistributedLedgerWithoutANativeDigitalTokenBlockchain DigitalTokenType = [2]int{2, 1}
	DistributedLedgerWithoutANativeDigitalTokenOther      DigitalTokenType = [2]int{2, 0}
	FunctionallyFungibleGroupOfDigitalTokens              DigitalTokenType = [2]int{3, 0}
)

// New creates a new DTIF instance
func New(log logger.Logger, distributedLedger, username, password string) *Dtif {
	return NewWithSourceAddress(log, distributedLedger, "https://download.dtif.org/data.json", username, password)
}

// NewWithSourceAddress creates a new DTIF instance from the requested source address
func NewWithSourceAddress(log logger.Logger, distributedLedger, sourceAddress, username, password string) *Dtif {
	return &Dtif{
		log,
		distributedLedger,
		sourceAddress,
		username,
		password,
		make(map[string]string),
		make(map[string]*big.Int),
		make(map[string]string),
		"",
		sync.RWMutex{},
	}
}

// Update fetches latest entries from the source, whether an url or a file
func (d *Dtif) Update(ctx context.Context) error {
	addrUrl, err := url.Parse(d.sourceAddress)
	if err != nil {
		return err
	}

	var content []byte

	if addrUrl.Scheme == "file" {
		filePath := strings.ReplaceAll(d.sourceAddress, "file://", "")

		stat, err := os.Stat(filePath)
		if err != nil {
			return err
		}

		newVersion := stat.ModTime().String()
		if newVersion == d.lastVersion {
			d.log.Debug(ctx, "DTIF data is not changed, no need update")
			return nil
		}

		content, err = os.ReadFile(filePath)
		if err != nil {
			return err
		}

		d.lastVersion = stat.ModTime().String()
		d.log.Debug(ctx, "DTIF data updated")
	} else {
		accessToken, err := d.Login(ctx, "masih@coreum.com", "B$cw12AV%O3NS4w$")
		if err != nil {
			return err
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, d.sourceAddress, nil)
		if err != nil {
			return err
		}
		req.Header.Set("authorization", "Bearer "+accessToken)
		req.Header.Set("If-None-Match", d.lastVersion)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(res.Body)

		if res.StatusCode >= http.StatusBadRequest {
			return fmt.Errorf("status %d: %s", res.StatusCode, res.Status)
		}

		newVersion := res.Header.Get("ETag")
		if newVersion == d.lastVersion {
			d.log.Debug(ctx, "DTIF data is not changed, no need update")
			return nil
		}

		content, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		d.lastVersion = newVersion
		d.log.Debug(ctx, "DTIF data updated")
	}

	var data map[string][]json.RawMessage
	err = json.Unmarshal(content, &data)
	if err != nil {
		return err
	}

	dtiToDenom := make(map[string]string)
	dtiToPriceMultiplier := make(map[string]*big.Int)
	denomToDti := make(map[string]string)

	for _, item := range data["records"] {
		var temp Record
		err = json.Unmarshal(item, &temp)
		if err != nil {
			return err
		}
		if temp.Header.DTIType != 0 || temp.Header.DLTType != 0 {
			continue
		}

		var record DigitalToken = new(AuxiliaryDigitalTokenJson)
		err = json.Unmarshal(item, record)
		if err != nil {
			return err
		}

		token, ok := record.(*AuxiliaryDigitalTokenJson)
		if !ok {
			// we are only interested in AuxiliaryDigitalTokens
			continue
		}

		if token.Normative.AuxiliaryDistributedLedger == "<locked>" || token.Normative.AuxiliaryDistributedLedger != d.distributedLedger {
			// we are only interested in Coreum tokens
			continue
		}

		denom := record.Denom()
		if denom == nil {
			return fmt.Errorf("token %s has no denom", record.DTI()) // TODO
		}

		dtiToDenom[record.DTI()] = *denom
		dtiToPriceMultiplier[record.DTI()] = record.PriceMultiplier()
		denomToDti[*denom] = record.DTI()
	}

	d.lock.Lock()
	d.dtiToDenom = dtiToDenom
	d.dtiToPriceMultiplier = dtiToPriceMultiplier
	d.denomToDti = denomToDti
	d.lock.Unlock()
	return nil
}

// LookupByDTI tries to find a specific token denom and price multiplier using its DTI
func (d *Dtif) LookupByDTI(dti string) (string, *big.Int, bool) {
	// It is also possible to get information of a token from:
	// https://download.dtif.org/Tokens/{dti}/Record/{dti}.json
	d.lock.RLock()
	defer d.lock.RUnlock()
	denom, found := d.dtiToDenom[dti]
	if found {
		return denom, d.dtiToPriceMultiplier[dti], true
	}
	return denom, nil, found
}

// LookupByDenom tries to find a specific token DTI using its denom
func (d *Dtif) LookupByDenom(denom string) (string, bool) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	dti, found := d.denomToDti[denom]
	return dti, found
}

var embeddedConfigPattern = regexp.MustCompile("atob\\('([^']+)'\\)")
var formFieldsPattern = regexp.MustCompile("name=\"([^\"]+)\"\\s+value=\"([^\"]+)\"")

func (d *Dtif) Login(ctx context.Context, username, password string) (string, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{
		Transport: http.DefaultTransport,
		Jar:       jar,
	}

	res, err := client.Get("https://auth.dtif.org/authorize?response_type=code&client_id=eaYmwHhnZUgMhPUBuhC9GV867aqKDqon&connection=DB-DTIF&redirect_uri=https://dtif.org/wp-login.php?redirect_to=https%3A%2F%2Fdtif.org%2F&reauth=1")
	if err != nil {
		return "", err
	}

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("dtif status %d", res.StatusCode)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	parts := embeddedConfigPattern.FindStringSubmatch(string(body))

	if len(parts) < 2 {
		return "", errors.New("DTIF implementation changed")
	}

	jsonBytes, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	var loginConfig LoginConfig
	err = json.Unmarshal(jsonBytes, &loginConfig)
	if err != nil {
		return "", err
	}

	payloadData := LoginPayload{
		ClientId:     loginConfig.ClientID,
		RedirectUri:  loginConfig.CallbackURL,
		Tenant:       loginConfig.Auth0Tenant,
		ResponseType: "token id_token",
		Csrf:         loginConfig.InternalOptions.Csrf,
		State:        loginConfig.InternalOptions.State,
		Intstate:     loginConfig.InternalOptions.Intstate,
		Username:     username,
		Password:     password,
		Nonce:        "rC8xYDMd1ncElGwXCSRKzbvRvV52V1n9",
		Connection:   loginConfig.Connection,
	}

	payload, err := json.Marshal(payloadData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://auth.dtif.org/usernamepassword/login", bytes.NewReader(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("auth0-client", "eyJuYW1lIjoiYXV0aDAuanMtdWxwIiwidmVyc2lvbiI6IjkuMTEuMiJ9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")

	res, err = client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("dtif status %d", res.StatusCode)
	}

	defer res.Body.Close()
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	fields := formFieldsPattern.FindAllStringSubmatch(string(body), -1)

	formPayload := url.Values{}
	for _, field := range fields {
		formPayload.Set(field[1], html.UnescapeString(field[2]))
	}

	req, err = http.NewRequestWithContext(ctx, http.MethodPost, "https://auth.dtif.org/login/callback", strings.NewReader(formPayload.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")

	res, err = client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("dtif status %d", res.StatusCode)
	}

	defer res.Body.Close()
	cookies := res.Request.Cookies()
	accessToken := ""
	for _, cookie := range cookies {
		if cookie.Name == "access_token" {
			accessToken = cookie.Value
		}
	}

	return accessToken, nil
}

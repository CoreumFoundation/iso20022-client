package dtif

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/pkg/errors"

	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
)

type Dtif struct {
	log               logger.Logger
	distributedLedger string
	sourceAddress     string
	dtiToDenom        map[string]string
	denomToDti        map[string]string
	lastVersion       string
	lock              sync.RWMutex
}

type DigitalToken interface {
	DTI() string
	Denom() *string
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
func New(log logger.Logger, distributedLedger string) *Dtif {
	return NewWithSourceAddress(log, distributedLedger, "https://download.dtif.org/data.json")
}

// NewWithSourceAddress creates a new DTIF instance from the requested source address
func NewWithSourceAddress(log logger.Logger, distributedLedger, sourceAddress string) *Dtif {
	return &Dtif{
		log,
		distributedLedger,
		sourceAddress,
		make(map[string]string),
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
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, d.sourceAddress, nil)
		if err != nil {
			return err
		}
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
	denomToDti := make(map[string]string)

	for _, item := range data["records"] {
		var temp Record
		err = json.Unmarshal(item, &temp)
		if err != nil {
			return err
		}
		ty := [2]int{temp.Header.DTIType, temp.Header.DLTType}
		var record DigitalToken
		switch ty {
		case AuxiliaryDigitalToken:
			record = new(AuxiliaryDigitalTokenJson)
		case NativeDigitalTokenBlockchain:
			record = new(NativeDigitalTokenBlockchainJson)
		case NativeDigitalTokenOther:
			record = new(NativeDigitalTokenOtherJson)
		case DistributedLedgerWithoutANativeDigitalTokenBlockchain:
			record = new(DistributedLedgerWithoutANativeDigitalTokenBlockchainJson)
		case DistributedLedgerWithoutANativeDigitalTokenOther:
			record = new(DistributedLedgerWithoutANativeDigitalTokenOtherJson)
		case FunctionallyFungibleGroupOfDigitalTokens:
			record = new(FunctionallyFungibleGroupOfDigitalTokensJson)
		default:
			return errors.New("unsupported token type") // TODO
		}
		err = json.Unmarshal(item, record)
		if err != nil {
			return err
		}

		token, ok := record.(*AuxiliaryDigitalTokenJson)
		if !ok {
			// we are only interested in AuxiliaryDigitalTokens
			continue
		}

		if token.Normative.AuxiliaryDistributedLedger == nil || *token.Normative.AuxiliaryDistributedLedger != d.distributedLedger {
			// we are only interested in Coreum tokens
			continue
		}

		denom := record.Denom()
		if denom == nil {
			return fmt.Errorf("token %s has no denom", record.DTI()) // TODO
		}

		dtiToDenom[record.DTI()] = *denom
		denomToDti[*denom] = record.DTI()
	}

	d.lock.Lock()
	d.dtiToDenom = dtiToDenom
	d.denomToDti = denomToDti
	d.lock.Unlock()
	return nil
}

// LookupByDTI tries to find a specific token denom using its DTI
func (d *Dtif) LookupByDTI(dti string) (string, bool) {
	// It is also possible to get information of a token from:
	// https://download.dtif.org/Tokens/{dti}/Record/{dti}.json
	d.lock.RLock()
	defer d.lock.RUnlock()
	denom, found := d.dtiToDenom[dti]
	return denom, found
}

// LookupByDenom tries to find a specific token DTI using its denom
func (d *Dtif) LookupByDenom(denom string) (string, bool) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	dti, found := d.denomToDti[denom]
	return dti, found
}

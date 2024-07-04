package addressbook

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

type AddressBook struct {
	repoAddress       string
	storedAddressBook StoredAddressBook
	lastVersion       string
	lock              sync.RWMutex
}

// New creates a new address book
func New(chainId string) *AddressBook {
	// TODO: replace with main branch after release
	repo := "https://raw.githubusercontent.com/CoreumFoundation/iso20022-addressbook/develop/%s/addressbook.json"
	return NewWithRepoAddress(fmt.Sprintf(repo, chainId))
}

// NewWithRepoAddress creates a new address book from requested repo address
func NewWithRepoAddress(repoAddress string) *AddressBook {
	return &AddressBook{
		repoAddress,
		StoredAddressBook{},
		"",
		sync.RWMutex{},
	}
}

// Update fetches latest entries from the repo, whether an url or a file
func (a *AddressBook) Update(ctx context.Context) error {
	addrUrl, err := url.Parse(a.repoAddress)
	if err != nil {
		return err
	}

	var content []byte

	if addrUrl.Scheme == "file" {
		filePath := strings.ReplaceAll(a.repoAddress, "file://", "")

		stat, err := os.Stat(filePath)
		if err != nil {
			return err
		}

		newVersion := stat.ModTime().String()
		if newVersion == a.lastVersion {
			return nil
		}

		content, err = os.ReadFile(filePath)
		if err != nil {
			return err
		}

		a.lastVersion = stat.ModTime().String()
	} else {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.repoAddress, nil)
		if err != nil {
			return err
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(res.Body)

		if res.StatusCode >= 400 {
			return fmt.Errorf("status %d: %s", res.StatusCode, res.Status)
		}

		newVersion := res.Header.Get("ETag")
		if newVersion == a.lastVersion {
			return nil
		}

		content, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		a.lastVersion = newVersion
	}

	var response StoredAddressBook
	err = json.Unmarshal(content, &response)
	if err != nil {
		return err
	}

	a.lock.Lock()
	a.storedAddressBook = response
	a.lock.Unlock()
	return nil
}

func (a *AddressBook) Validate() error {
	localAddressBook := make(map[string]struct{})

	a.lock.RLock()
	defer a.lock.RUnlock()

	for _, address := range a.storedAddressBook.Addresses {
		if _, alreadyExists := localAddressBook[address.Bech32EncodedAddress]; alreadyExists {
			return fmt.Errorf(
				"duplicate entries with bech32 encoded address %q",
				address.Bech32EncodedAddress,
			)
		}

		publicKeyBytes, err := base64.StdEncoding.DecodeString(address.PublicKey)
		if err != nil {
			return fmt.Errorf(
				"public key of %q is not a valid base64 encoded string: %v",
				address.Bech32EncodedAddress,
				err,
			)
		}

		if _, err = secp256k1.ParsePubKey(publicKeyBytes); err != nil {
			return fmt.Errorf(
				"public key of %q is not a valid secp256k1 public key: %v",
				address.Bech32EncodedAddress,
				err,
			)
		}

		localAddressBook[address.Bech32EncodedAddress] = struct{}{}
		matches := make([]string, 0, 1)

		for _, address2 := range a.storedAddressBook.Addresses {
			if address.BranchAndIdentification.Equal(address2.BranchAndIdentification) {
				matches = append(matches, address2.Bech32EncodedAddress)
				if len(matches) > 1 {
					break
				}
			}
		}

		if len(matches) > 1 {
			return fmt.Errorf(
				"ISO20022 branch and identification data of entry %q and %q conflicts",
				matches[0],
				matches[1],
			)
		}
	}

	return nil
}

// Lookup tries to find a specific entry in the address book using ISO20022 BranchAndIdentification data
func (a *AddressBook) Lookup(expectedAddress BranchAndIdentification) (*Address, bool) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	for _, lookedUpAddress := range a.storedAddressBook.Addresses {
		if lookedUpAddress.BranchAndIdentification.Equal(expectedAddress) {
			return &lookedUpAddress, true
		}
	}
	return nil, false
}

// LookupByAccountAddress tries to find a specific entry in the address book using bech32-encoded account address
func (a *AddressBook) LookupByAccountAddress(bech32EncodedAddress string) (*Address, bool) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	for _, lookedUpAddress := range a.storedAddressBook.Addresses {
		if lookedUpAddress.Bech32EncodedAddress == bech32EncodedAddress {
			return &lookedUpAddress, true
		}
	}
	return nil, false
}

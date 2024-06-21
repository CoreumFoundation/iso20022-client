package addressbook

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

type AddressBook struct {
	repoAddress       string
	networkType       string
	storedAddressBook StoredAddressBook
	lock              sync.RWMutex
}

// New creates a new address book and fetch latest update entries
func New(networkType string) (*AddressBook, error) {
	addressBook := &AddressBook{
		"https://github.com/CoreumFoundation/iso20022-addressbook/raw/develop/%s/addressbook.json",
		networkType,
		StoredAddressBook{},
		sync.RWMutex{},
	}
	err := addressBook.Update()
	if err != nil {
		return nil, err
	}
	return addressBook, nil
}

// NewWithRepoAddress creates a new address book and fetch latest update entries from requested repo address
func NewWithRepoAddress(networkType, repoAddress string) (*AddressBook, error) {
	addressBook := &AddressBook{
		repoAddress,
		networkType,
		StoredAddressBook{},
		sync.RWMutex{},
	}
	err := addressBook.Update()
	if err != nil {
		return nil, err
	}
	return addressBook, nil
}

// Update fetches latest entries from the repo, whether an url or a file
func (a *AddressBook) Update() error {
	addr := fmt.Sprintf(a.repoAddress, a.networkType)

	addrUrl, err := url.Parse(addr)
	if err != nil {
		return err
	}

	var content []byte

	if addrUrl.Scheme == "file" {
		content, err = os.ReadFile(strings.ReplaceAll("file://", "", addr))
	} else {
		res, err := http.Get(addr)
		if err != nil {
			return err
		}

		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(res.Body)

		if res.StatusCode >= 400 {
			return fmt.Errorf("status %d: %s", res.StatusCode, res.Status)
		}

		content, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
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

// KeyAlgo returns keys algorithm for this chain
func (a *AddressBook) KeyAlgo() string {
	return a.storedAddressBook.KeyAlgo
}

// ForEach loops through add address book entries
func (a *AddressBook) ForEach(f func(address Address) bool) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	for _, addr := range a.storedAddressBook.Addresses {
		if f(addr) == false {
			return
		}
	}
}

// Lookup tries to find a specific entry in the address book using ISO20022 BranchAndIdentification data
func (a *AddressBook) Lookup(wantedAddress BranchAndIdentification) (*Address, bool) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	for _, lookedUpAddress := range a.storedAddressBook.Addresses {
		lookedUpId := lookedUpAddress.BranchAndIdentification.Identification
		wantedId := wantedAddress.Identification

		if lookedUpId.Bic == wantedId.Bic || lookedUpId.Lei == wantedId.Lei {
			return &lookedUpAddress, true
		}

		if lookedUpId.ClearingSystemMemberIdentification != nil && wantedId.ClearingSystemMemberIdentification != nil {
			lookedUpCls := lookedUpId.ClearingSystemMemberIdentification
			wantedCls := wantedId.ClearingSystemMemberIdentification
			if lookedUpCls.MemberId == wantedCls.MemberId {
				return &lookedUpAddress, true
			}
			if lookedUpCls.ClearingSystemId != nil && wantedCls.ClearingSystemId != nil {
				if lookedUpCls.ClearingSystemId.String() == wantedCls.ClearingSystemId.String() {
					return &lookedUpAddress, true
				}
			}
		}

		if lookedUpId.PostalAddress != nil && wantedId.PostalAddress != nil {
			if lookedUpId.Name == wantedId.Name && lookedUpId.PostalAddress.String() == wantedId.PostalAddress.String() {
				return &lookedUpAddress, true
			}
		}

		if (lookedUpId.Other != nil && wantedId.Other != nil) && lookedUpId.Other.String() == wantedId.Other.String() {
			return &lookedUpAddress, true
		}
	}
	return nil, false
}

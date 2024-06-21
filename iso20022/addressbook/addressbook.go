package addressbook

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type AddressBook struct {
	repoAddress       string
	networkType       string
	storedAddressBook StoredAddressBook
	lock              sync.RWMutex
}

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

func (a *AddressBook) Update() error {
	addr := fmt.Sprintf(a.repoAddress, a.networkType)
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var response StoredAddressBook
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	a.lock.Lock()
	a.storedAddressBook = response
	a.lock.Unlock()
	return nil
}

func (a *AddressBook) KeyAlgo() string {
	return a.storedAddressBook.KeyAlgo
}

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

package addressbook

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
)

type AddressBook struct {
	repoAddress       string
	chainId           string
	storedAddressBook StoredAddressBook
	lastVersion       string
	lock              sync.RWMutex
}

// New creates a new address book
func New(chainId string) *AddressBook {
	// TODO: replace with main branch after release
	repo := "https://raw.githubusercontent.com/CoreumFoundation/iso20022-addressbook/develop/%s/addressbook.json"
	return NewWithRepoAddress(chainId, repo)
}

// NewWithRepoAddress creates a new address book from requested repo address
func NewWithRepoAddress(chainId, repoAddress string) *AddressBook {
	return &AddressBook{
		repoAddress,
		chainId,
		StoredAddressBook{},
		"",
		sync.RWMutex{},
	}
}

// Update fetches latest entries from the repo, whether an url or a file
func (a *AddressBook) Update(ctx context.Context) error {
	addr := fmt.Sprintf(a.repoAddress, a.chainId)

	addrUrl, err := url.Parse(addr)
	if err != nil {
		return err
	}

	var content []byte

	if addrUrl.Scheme == "file" {
		filePath := strings.ReplaceAll(addr, "file://", "")

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
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, addr, nil)
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

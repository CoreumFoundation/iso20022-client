package processes

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
)

// AddressBookUpdaterProcess is the process that keeps the address book updated.
type AddressBookUpdaterProcess struct {
	interval    time.Duration
	log         logger.Logger
	addressBook AddressBook
}

// NewAddressBookUpdaterProcess returns a new instance of the AddressBookUpdaterProcess.
func NewAddressBookUpdaterProcess(
	interval time.Duration,
	log logger.Logger,
	addressBook AddressBook,
) (*AddressBookUpdaterProcess, error) {
	if interval < time.Second {
		return nil, errors.Errorf("failed to init the process, interval cannot be less than one second")
	}

	return &AddressBookUpdaterProcess{
		interval,
		log,
		addressBook,
	}, nil
}

// Start starts the process.
func (p *AddressBookUpdaterProcess) Start(ctx context.Context) error {
	p.log.Info(ctx, "Starting the address book updater process")

	ticker := time.NewTicker(p.interval)
	for {
		select {
		case <-ticker.C:
			if err := p.addressBook.Update(ctx); err != nil {
				return err
			}
		case <-ctx.Done():
			ticker.Stop()
			return nil
		}
	}
}

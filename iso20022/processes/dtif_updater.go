package processes

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
)

// DtifUpdaterProcess is the process that keeps the DTIF updated.
type DtifUpdaterProcess struct {
	interval time.Duration
	log      logger.Logger
	dtif     Dtif
}

// NewDtifUpdaterProcess returns a new instance of the DtifUpdaterProcess.
func NewDtifUpdaterProcess(
	interval time.Duration,
	log logger.Logger,
	dtif Dtif,
) (*DtifUpdaterProcess, error) {
	if interval < time.Second {
		return nil, errors.Errorf("failed to init the process, interval cannot be less than one second")
	}

	return &DtifUpdaterProcess{
		interval,
		log,
		dtif,
	}, nil
}

// Start starts the process.
func (p *DtifUpdaterProcess) Start(ctx context.Context) error {
	p.log.Info(ctx, "Starting the DTIF updater process")

	ticker := time.NewTicker(p.interval)
	for {
		select {
		case <-ticker.C:
			if err := p.dtif.Update(ctx); err != nil {
				return err
			}
		case <-ctx.Done():
			ticker.Stop()
			return nil
		}
	}
}

package services

import (
	"context"

	"github.com/gorcz/ports/internal/datastore"
	"github.com/gorcz/ports/pkg/model/port"
)

// nolint: gocritic,lll
//go:generate mockgen -destination=./../../mocks/pkg/services/ports.go -package=mocks github.com/gorcz/ports/internal/services Ports

type Ports interface {
	UpsertPorts(ctx context.Context, portIterator port.Iterator) error
}

type Service struct {
	datastore datastore.Datastore
}

func NewPorts(datastore datastore.Datastore) *Service {
	return &Service{
		datastore: datastore,
	}
}

func (ps *Service) UpsertPorts(ctx context.Context, portIterator port.Iterator) error {
	if portIterator == nil {
		return nil
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			portData, ok, err := portIterator.Next()
			if err != nil {
				return err
			}
			if !ok {
				return nil
			}
			if err = ps.upsertPort(ctx, portData); err != nil {
				return err
			}
		}
	}
}

func (ps *Service) upsertPort(ctx context.Context, port *port.Port) error {
	return ps.datastore.UpsertPort(ctx, port)
}

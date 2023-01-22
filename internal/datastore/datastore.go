package datastore

import (
	"context"

	"github.com/gorcz/ports/pkg/model/port"
)

// nolint: gocritic,lll
//go:generate mockgen -destination=./../../mocks/pkg/datastore/datastore.go -package=mocks github.com/gorcz/ports/internal/datastore Datastore

type Datastore interface {
	UpsertPort(ctx context.Context, port *port.Port) error
}

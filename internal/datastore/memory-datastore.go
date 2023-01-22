package datastore

import (
	"context"
	"sync"

	"github.com/gorcz/ports/pkg/model/port"
)

type MemoryDatastore struct {
	sync.RWMutex

	ports map[port.Code]port.Details
}

func NewMemoryDatastore() *MemoryDatastore {
	return &MemoryDatastore{
		ports: make(map[port.Code]port.Details),
	}
}

func (ds *MemoryDatastore) UpsertPort(_ context.Context, port *port.Port) error {
	if port == nil {
		return nil
	}

	ds.Lock()
	defer ds.Unlock()

	ds.ports[port.Code] = port.Details
	return nil
}

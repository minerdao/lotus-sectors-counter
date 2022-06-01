package storedcounter

import (
	"context"
	"encoding/binary"
	"sync"

	"github.com/ipfs/go-datastore"
	"golang.org/x/xerrors"
)

// StoredCounter is a counter that persists to a datastore as it increments
type StoredCounter struct {
	lock sync.Mutex
	ds   datastore.Datastore
	name datastore.Key
}

// New returns a new StoredCounter for the given datastore and key
func New(ds datastore.Datastore, name datastore.Key) *StoredCounter {
	return &StoredCounter{ds: ds, name: name}
}

// Next returns the next counter value, updating it on disk in the process
// if no counter is present, it creates one and returns a 0 value
func (sc *StoredCounter) Next() (uint64, error) {
	ctx := context.TODO()
	sc.lock.Lock()
	defer sc.lock.Unlock()

	has, err := sc.ds.Has(ctx, sc.name)
	if err != nil {
		return 0, err
	}

	var next uint64 = 0
	if has {
		curBytes, err := sc.ds.Get(ctx, sc.name)
		if err != nil {
			return 0, err
		}
		cur, _ := binary.Uvarint(curBytes)
		next = cur + 1
	}
	buf := make([]byte, binary.MaxVarintLen64)
	size := binary.PutUvarint(buf, next)

	return next, sc.ds.Put(ctx, sc.name, buf[:size])
}

// Get current Sector Number
func (sc *StoredCounter) Get() (uint64, error) {
	ctx := context.TODO()
	sc.lock.Lock()
	defer sc.lock.Unlock()
	has, err := sc.ds.Has(ctx, sc.name)
	if err != nil {
		return 0, err
	}

	if ! has {
		return 0, nil
	}

	curBytes, err := sc.ds.Get(ctx, sc.name)
	if err != nil {
		return 0, err
	}
	cur, _ := binary.Uvarint(curBytes);
	return cur,nil
}

// Set the next counter value, updating it on disk in the process
func (sc *StoredCounter) Set(number uint64) error {
	ctx := context.TODO()
	sc.lock.Lock()
	defer sc.lock.Unlock()
	has, err := sc.ds.Has(ctx, sc.name)
	if err != nil {
		return err
	}

	if has {
		curBytes, err := sc.ds.Get(ctx, sc.name)
		if err != nil {
			return err
		}
		cur, _ := binary.Uvarint(curBytes)
		if cur > number {
			return xerrors.Errorf("Number %d should not less than current SectorID %d", number, cur)
		}
	}

	buf := make([]byte, binary.MaxVarintLen64)
	size := binary.PutUvarint(buf, number)
	return sc.ds.Put(ctx, sc.name, buf[:size])
}

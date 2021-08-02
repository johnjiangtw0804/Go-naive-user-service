package models

import "golang.org/x/mod/sumdb/storage"

// StorageType defines available storage types
type StorageType int

const (
	// Memory will store data in memory
	Memory = iota
)

var DB Storage

type Storage interface {
}

func NewStorage(storageType StorageType) error {
	var err error

	switch storageType {
	case Memory:
		DB, err = new(storage.Memory)
		if err != nil {
			return err
		}
	}

	return nil
}

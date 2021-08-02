package models

import "golang.org/x/mod/sumdb/storage"

// StorageType variable type defines available storage types
type StorageType int

const (
	// StorageType option
	Memory = iota
)

var DB Storage

type Storage interface {
	GetUser() error
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

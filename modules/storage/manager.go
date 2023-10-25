package storage

import "errors"

type StorageManager struct {
	store Store
}

func (manager StorageManager) GetStore() Store {
	return manager.store
}

var manager *StorageManager

func InitManager(store Store) {
	manager = &StorageManager{
		store: store,
	}
}

func GetManager() (*StorageManager, error) {
	if manager == nil {
		return nil, errors.New("Manager not initialized")
	}

	return manager, nil
}

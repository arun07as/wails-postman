package stores

import (
	"errors"

	"github.com/arun07as/wails-postman/modules/storage"
)

// TODO Implement filesystem operations

type Filesystem struct {
}

func (f Filesystem) Get(entity storage.Entity, filters storage.Filters) (*storage.Result, error) {
	return nil, errors.New("Not Implemented")
}

func (f Filesystem) Save(entity storage.Entity, inout storage.Input) error {
	return errors.New("Not Implemented")
}

func (f Filesystem) Delete(entity storage.Entity, inout storage.Input) error {
	return errors.New("Not Implemented")
}

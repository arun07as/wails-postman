package storage

type Store interface {
	Get(entity Entity, filters Filters) (*Result, error)
	Save(entity Entity, input Input) error
	Delete(entity Entity, input Input) error
}

package storage

type Entity interface {
	GetName() string
}

type Filter struct {
	Field     string
	Operation string
	Value     string
}

type Filters interface {
	GetFilters() []Filter
}

type Input struct {
	input interface{}
}

func (input *Input) GetInput() interface{} {
	return input.input
}

type Result struct {
	result interface{}
}

func (result *Result) GetResult() interface{} {
	return result.result
}

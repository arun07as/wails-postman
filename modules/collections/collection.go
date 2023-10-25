package collections

import (
	"github.com/arun07as/wails-postman/modules/environments"
	"github.com/arun07as/wails-postman/modules/generic"
)

type Collection struct {
	generic.ItemInfo
	Variables    []environments.EnvironmentItem `json:"variables"`
	Environments []environments.Environment     `json:"environments"`
	Environment  environments.Environment       `json:"environment"`
	Items        []CollectionItem               `json:"items"`
}

func (c Collection) GetVariables() []environments.EnvironmentItem {
	return c.Variables
}

func (c Collection) GetEnvironments() []environments.Environment {
	return c.Environments
}

func (c Collection) GetEnvironment() environments.Environment {
	return c.Environment
}

func (c Collection) GetItems() []CollectionItem {
	return c.Items
}

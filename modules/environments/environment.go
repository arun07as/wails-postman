package environments

import (
	"errors"

	"github.com/arun07as/wails-postman/modules/generic"
	"golang.org/x/exp/slices"
)

type EnvironmentItem struct {
	generic.Variable
	InitialValue string `json:"initial_value"`
}

func (item EnvironmentItem) GetInitialValue() string {
	return item.InitialValue
}

type Environment struct {
	generic.ItemInfo
	Items []EnvironmentItem `json:"items"`
}

func (env Environment) GetItems() []EnvironmentItem {
	return env.Items
}

func (env *Environment) AddItem(item EnvironmentItem) error {
	env.Items = append(env.GetItems(), item)
	return nil
}

func (env *Environment) RemoveItem(item EnvironmentItem) error {
	index := slices.Index(env.Items, item)

	if index == -1 {
		return errors.New("Cannot find item")
	}

	env.Items = slices.Delete(env.Items, index, index)

	return nil
}

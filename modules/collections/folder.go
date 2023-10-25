package collections

import "github.com/arun07as/wails-postman/modules/generic"

type Folder struct {
	generic.ItemInfo
	Items []CollectionItem `json:"items"`
}

func (f Folder) GetItems() []CollectionItem {
	return f.Items
}

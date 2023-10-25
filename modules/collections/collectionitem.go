package collections

type CollectionItem struct {
	IsAFolder bool    `json:"is_folder"`
	Folder    Folder  `json:"folder,omitempty"`
	Request   Request `json:"request,omitempty"`
}

func (item CollectionItem) IsFolder() bool {
	return item.IsAFolder
}

func (item CollectionItem) GetFolder() Folder {
	return item.Folder
}

func (item CollectionItem) GetRequest() Request {
	return item.Request
}

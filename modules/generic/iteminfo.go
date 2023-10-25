package generic

type ItemInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (info ItemInfo) GetName() string {
	return info.Name
}

func (info ItemInfo) GetDescription() string {
	return info.Description
}

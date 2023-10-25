package generic

type Variable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (v Variable) GetName() string {
	return v.Name
}

func (v Variable) GetValue() string {
	return v.Value
}

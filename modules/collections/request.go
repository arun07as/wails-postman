package collections

import "github.com/arun07as/wails-postman/modules/generic"

type Request struct {
	generic.ItemInfo
	Url         string `json:"url"`
	Method      string `json:"method"`       // TODO replace string with some enum implementation
	Body        string `json:"body"`         // TODO replace string with proper body struct
	QueryParams string `json:"query_params"` // TODO replace string with proper query params struct
	PathParams  string `json:"path_params"`  // TODO replace string with proper path params struct
}

func (r Request) GetUrl() string {
	return r.Url
}

func (r Request) GetMethod() string {
	return r.Method
}

func (r Request) GetBody() string {
	return r.Body
}

func (r Request) GetQueryParams() string {
	return r.QueryParams
}

func (r Request) GetPathParams() string {
	return r.PathParams
}

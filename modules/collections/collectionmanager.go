package collections

import (
	"github.com/arun07as/wails-postman/modules/environments"
	"github.com/arun07as/wails-postman/modules/generic"
)

type CollectionManager struct {
}

func (manager CollectionManager) GetCollections() []Collection {

	environment := environments.Environment{
		ItemInfo: generic.ItemInfo{
			Name:        "Environment Name",
			Description: "Environment Description",
		},
		Items: []environments.EnvironmentItem{
			{
				Variable: generic.Variable{
					Name:  "Test Environment Name",
					Value: "Test Environment Value",
				},
				InitialValue: "Test Environment InitialValue",
			},
		},
	}

	collection := Collection{
		ItemInfo: generic.ItemInfo{
			Name:        "Test",
			Description: "Description Test",
		},
		Variables: []environments.EnvironmentItem{
			{
				Variable: generic.Variable{
					Name:  "Test Variable Name",
					Value: "Test Variable Value",
				},
				InitialValue: "Test Initial Variable Value",
			},
		},
		Environments: []environments.Environment{
			environment,
		},
		Environment: environment,
		Items: []CollectionItem{
			{
				IsAFolder: true,
				Folder: Folder{
					ItemInfo: generic.ItemInfo{
						Name:        "Test Folder Name",
						Description: "Test Description",
					},
					Items: []CollectionItem{
						{
							IsAFolder: false,
							Request: Request{
								ItemInfo: generic.ItemInfo{
									Name:        "Test Folder Name",
									Description: "Test Description",
								},
								Url:         "https://test@example.com",
								Method:      "GET",
								Body:        "{\"hello\":\"World\"}",
								QueryParams: "",
								PathParams:  "",
							},
						},
					},
				},
			},
		},
	}
	return []Collection{
		collection,
	}
}

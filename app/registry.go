package app

import "github.com/wibu-elite/gotoko/app/models"

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Address{}},
		{Model: models.Categories{}},
		{Model: models.Product{}},
		{Model: models.ProductImages{}},
		{Model: models.Section{}},
	}
}

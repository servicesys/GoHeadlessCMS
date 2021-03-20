package storage

import (
	"github.com/servicesys/GoHeadlessCMS/pkg/model"
)

type ContentStorage interface {
	SearchContent(fields []string, query []string) ([]model.Content, error)
	GetContent(uuid string) (model.Content, error)
	GetContentByName(name string) (model.Content, error)
	GetAllContentByCategory(categoryCod string) ([]model.Content, error)
	CreateContent(model.Content) error
	UpdateContent(model.Content) error
	DeleteContent(uuid string) error
	CreateCategory(model.Category) error
	UpdateCategory(model.Category) error
	GetAllCategory() ([]model.Category, error)
	CreateType(model.Type) error
	UpdateType(model.Type) error
	GetType(cod string) (model.Type, error)
}

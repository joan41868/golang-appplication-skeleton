package repository

import "github.com/joan41868/go-app-skeleton/src/model"

// Repository defines generic operations on a database level.
// Each of those operations is applicable to many projects.
type Repository interface {
	Create(e *model.Entity) (*model.Entity, error)

	GetById(id interface{}) (*model.Entity, error)

	GetMany() (*model.Entity, error)

	Update(id interface{}, ne *model.Entity) (*model.Entity, error)

	Delete(id interface{}) error
}

package repository

import (
	"gitlab.com/imagyn_go/internal/domain/entity"
)

type Example interface {
	Create(example *entity.Example) (entity.Example, error)
	//Update(example *entity.Example) (*entity.Example, error)
	//FindById(id uuid.UUID) (*entity.Example, error)
	//FindAll() ([]*entity.Example, error)
	//Delete(id uuid.UUID) error
}

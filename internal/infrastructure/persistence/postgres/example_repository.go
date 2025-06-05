package postgres

import (
	"gitlab.com/imagyn_go/internal/domain/example"
	"gorm.io/gorm"
)

type ExampleRepository struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) example.Repository {
	return &ExampleRepository{db: db}
}

func (r *ExampleRepository) Create(example *example.Example) error {
	//TODO implement me
	return r.db.Create(example).Error
}

func (r *ExampleRepository) UpdateExample(example *example.Example) {
	//TODO implement me
	panic("implement me")
}

func (r *ExampleRepository) FindByName(name string) (*example.Example, error) {
	//TODO implement me
	var example example.Example
	err := r.db.Where("name = ?", name).First(&example).Error

	return &example, err
}

func (r *ExampleRepository) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}

package example

import (
	"errors"
	"time"
)

type Example struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(255);unique; not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewExample(name string) (*Example, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	return &Example{
		Name:      name,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}, nil
}

func (e *Example) UpdateName(name string) {
	e.Name = name
	e.UpdatedAt = time.Now()
}

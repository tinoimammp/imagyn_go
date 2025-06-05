package example

type Repository interface {
	Create(example *Example) error
	UpdateExample(example *Example)
	FindByName(name string) (*Example, error)
	Delete(id uint) error
}

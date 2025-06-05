package example

type ExampleService interface {
	Create(name string) (*Example, error)
}

type exampleService struct{}

func NewExampleService() ExampleService {
	return &exampleService{}
}

func (e exampleService) Create(name string) (*Example, error) {
	return NewExample(name)
}

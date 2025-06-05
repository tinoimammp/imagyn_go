package usecases

import (
	"gitlab.com/imagyn_go/internal/application/example/dto"
	"gitlab.com/imagyn_go/internal/domain/example"
)

type CreateUseCase struct {
	exampleRepo    example.Repository
	exampleService example.ExampleService
}

func NewExampleUseCase(repo example.Repository, service example.ExampleService) CreateUseCase {
	return CreateUseCase{
		exampleRepo:    repo,
		exampleService: service,
	}
}

func (uc *CreateUseCase) Excecute(req dto.CreateExampleRequest) error {
	example, err := uc.exampleService.Create(req.Name)
	if err != nil {
		return err
	}
	return uc.exampleRepo.Create(example)
}

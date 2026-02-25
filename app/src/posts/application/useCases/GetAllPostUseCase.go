package usecases

import (
	"sharedup/app/src/posts/domain/entities"
	"sharedup/app/src/posts/domain/repositories"
)

type GetAllPostUseCase struct {
	repo repositories.IPostRepository
}

func NewGetAllPostUseCase(repo repositories.IPostRepository) *GetAllPostUseCase {
	return &GetAllPostUseCase{repo: repo}
}

func (uc *GetAllPostUseCase) Execute() ([]entities.PostResponse, error) {
	return uc.repo.GetAll()
}
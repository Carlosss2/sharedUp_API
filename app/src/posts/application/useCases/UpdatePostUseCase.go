package usecases

import (
	"sharedup/app/src/posts/domain/entities"
	"sharedup/app/src/posts/domain/repositories"
)

type UpdatePostUseCase struct {
	repo repositories.IPostRepository
}

func NewUpdatePostUseCase(repo repositories.IPostRepository) *UpdatePostUseCase {
	return &UpdatePostUseCase{repo: repo}
}

func (uc *UpdatePostUseCase) Execute(p entities.Post) error {
	return uc.repo.Update(p)
}
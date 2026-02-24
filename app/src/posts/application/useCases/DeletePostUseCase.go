package usecases

import (
	"sharedup/app/src/posts/domain/repositories"
)

type DeletePostUseCase struct {
	repo repositories.IPostRepository
}

func NewDeletePostUseCase(repo repositories.IPostRepository) *DeletePostUseCase {
	return &DeletePostUseCase{repo: repo}
}

func (uc *DeletePostUseCase) Execute(id int, idUser int) error {
	return uc.repo.Delete(id, idUser)
}
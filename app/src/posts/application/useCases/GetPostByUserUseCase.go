package usecases

import (
	"sharedup/app/src/posts/domain/entities"
	"sharedup/app/src/posts/domain/repositories"
)

type GetPostByUserUseCase struct {
	repo repositories.IPostRepository
}

func NewGetPostByUserUseCase(repo repositories.IPostRepository) *GetPostByUserUseCase {
	return &GetPostByUserUseCase{repo: repo}
}

func (uc *GetPostByUserUseCase) Execute(userID int) ([]entities.PostResponse, error) {
	return uc.repo.GetByUser(userID)
}
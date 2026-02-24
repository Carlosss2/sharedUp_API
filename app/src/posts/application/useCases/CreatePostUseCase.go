package usecases

import (
	"sharedup/app/src/posts/domain/entities"
	"sharedup/app/src/posts/domain/repositories"
)

type CreatePostUseCase struct {
	repo repositories.IPostRepository
}

func NewCreatePostUseCase(repo repositories.IPostRepository)*CreatePostUseCase{
	return &CreatePostUseCase{repo: repo}
}
func(uc *CreatePostUseCase) Execute(p entities.Post) error{
 return uc.repo.Save(p)
}
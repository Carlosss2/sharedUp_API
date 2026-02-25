package usecases

import "sharedup/app/src/posts/domain/repositories"

type IncrementLikeUseCase struct {
	repo repositories.IPostRepository
}

func NewIncrementLikeUseCase(repo repositories.IPostRepository) *IncrementLikeUseCase {
	return &IncrementLikeUseCase{repo: repo}
}

func (uc *IncrementLikeUseCase) Execute(postID int) error {
	return uc.repo.IncrementLike(postID)
}
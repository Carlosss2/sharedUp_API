package usecases

import "sharedup/app/src/posts/domain/repositories"

type IncrementDislikeUseCase struct {
	repo repositories.IPostRepository
}

func NewIncrementDislikeUseCase(repo repositories.IPostRepository) *IncrementDislikeUseCase {
	return &IncrementDislikeUseCase{repo: repo}
}

func (uc *IncrementDislikeUseCase) Execute(postID int) error {
	return uc.repo.IncrementDislike(postID)
}
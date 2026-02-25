package usecases

import (
	"sharedup/app/src/posts/domain/entities"
	"sharedup/app/src/posts/domain/repositories"
)

type CreatePostUseCase struct {
	repo     repositories.IPostRepository
	notifier repositories.IPostNotifier
}

func NewCreatePostUseCase(
	repo repositories.IPostRepository,
	notifier repositories.IPostNotifier,
) *CreatePostUseCase {
	return &CreatePostUseCase{
		repo:     repo,
		notifier: notifier,
	}
}

func (uc *CreatePostUseCase) Execute(p entities.Post) error {

	err := uc.repo.Save(p)
	if err != nil {
		return err
	}

	// Notificar evento en tiempo real
	uc.notifier.NotifyPostCreated(p)

	return nil
}
package repositories

import "sharedup/app/src/posts/domain/entities"

type IPostNotifier interface {
	NotifyPostCreated(post entities.PostResponse)
}
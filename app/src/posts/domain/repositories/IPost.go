package repositories

import "sharedup/app/src/posts/domain/entities"

type IPostRepository interface {
	Save(post entities.Post) error
}
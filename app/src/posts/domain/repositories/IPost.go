package repositories

import "sharedup/app/src/posts/domain/entities"

type IPostRepository interface {
	Save(p entities.Post) (int64, error)
	GetByID(id int64) (entities.PostResponse, error)
	Update(p entities.Post) error
	Delete(id int, idUser int) error
	GetAll() ([]entities.PostResponse, error)
	GetByUser(idUser int) ([]entities.PostResponse, error)
	IncrementLike(postID int) error
	IncrementDislike(postID int) error
}
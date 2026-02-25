package repositories

import "sharedup/app/src/posts/domain/entities"

type IPostRepository interface {
	Save(post entities.Post) error
	Update (post entities.Post) error
	Delete(id int, idUser int) error
	GetAll() ([]entities.PostResponse, error)
	GetByUser(idUser int) ([]entities.PostResponse, error)
}
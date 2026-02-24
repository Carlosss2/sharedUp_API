package repositories

import "sharedup/app/src/register/domain/entities"

type IUserRepository interface {
	Save(user entities.User) error
}
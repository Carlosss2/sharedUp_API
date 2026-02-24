package repositories

import "sharedup/app/src/login/domain/entities"

type IloginRepository interface {
	LogIn(email string) (entities.UserWithPassword, error)
}
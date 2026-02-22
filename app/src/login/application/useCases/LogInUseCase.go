package usecases

import (
	"errors"
	"sharedup/app/src/login/domain/entities"
	"sharedup/app/src/login/domain/repositories"
	"sharedup/app/src/security/domain"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type LogInUseCase struct{
	db repositories.IloginRepository
	hash domain.PasswordHasher
}

func NewLogInUseCase(
	db repositories.IloginRepository,
	hash domain.PasswordHasher,
) *LogInUseCase {
	return &LogInUseCase{db: db, hash: hash}
}

func (uc *LogInUseCase) Execute(user entities.User) (entities.UserResponse, error) {
	dbUser, err := uc.db.LogIn(user.Email)
	if err != nil {
		return entities.UserResponse{}, ErrInvalidCredentials
	}

	if !uc.hash.Compare(dbUser.PasswordHash, user.Password) {
		return entities.UserResponse{}, ErrInvalidCredentials
	}

	return entities.UserResponse{
		Id:       dbUser.Id,
		Name:     dbUser.Name,
		Email:    dbUser.Email,
		Career:  dbUser.Career,
	}, nil
}

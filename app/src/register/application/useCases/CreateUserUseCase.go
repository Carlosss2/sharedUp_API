package usecases

import (
	"sharedup/app/src/register/domain/entities"
	"sharedup/app/src/register/domain/repositories"
	  hash "sharedup/app/src/security/infraestructure"
)

type CreateUserUseCase struct {
	db repositories.IUserRepository
}

func NewCreateUserUseCase(db repositories.IUserRepository)*CreateUserUseCase{
	return &CreateUserUseCase{db:db}
}

func (useCase *CreateUserUseCase) Execute(user entities.User) error {
	hashedPassword, err := hash.NewBCryptHasher().Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return useCase.db.Save(user)
}

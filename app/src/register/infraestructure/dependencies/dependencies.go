package dependencies

import (
	"sharedup/app/src/helpers"
	useCase "sharedup/app/src/register/application/useCases"
	"sharedup/app/src/register/infraestructure"
	"sharedup/app/src/register/infraestructure/controllers"
)

var (
	mySQL infraestructure.MySQL
	
)

func InitDependencies() {
	db, err := helpers.ConectToMySQL()
	if err != nil {
		panic(err)
	}

	mySQL = *infraestructure.NewMySQL(db)
}
func GetCreateUserController() *controllers.CreateUserController{
	createUserUseCase := useCase.NewCreateUserUseCase(&mySQL)
	return controllers.NewCreateUserController(createUserUseCase)
}
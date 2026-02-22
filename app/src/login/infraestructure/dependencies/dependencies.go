package dependencies

import (
	"sharedup/app/src/helpers"
	usecases "sharedup/app/src/login/application/useCases"
	"sharedup/app/src/login/infraestructure"
	"sharedup/app/src/login/infraestructure/controllers"
	infrastructureSecurity "sharedup/app/src/security/infraestructure"
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

func GetLogInController() *controllers.LogInController{
	hasher := infrastructureSecurity.NewBCryptHasher()
	useCase := usecases.NewLogInUseCase(&mySQL, hasher)
	return controllers.NewLogInController(useCase)
}
package dependencies

import (
	"sharedup/app/src/helpers"
	usecases "sharedup/app/src/posts/application/useCases"
	"sharedup/app/src/posts/infraestructure"
	"sharedup/app/src/posts/infraestructure/controllers"
)

var (
	mySQL infraestructure.MySQL
)

func InitPost(){
	db,err := helpers.ConectToMySQL()
	if err != nil {
		panic(err)
	}

	mySQL = *infraestructure.NewMySQL(db)
}

func GetCreatePostController() *controllers.CreatePostController{
	caseCreate := usecases.NewCreatePostUseCase(&mySQL)
	return controllers.NewCreatePostController(caseCreate)
}

// nuevas implementaciones el update y el delete

func GetUpdatePostController() *controllers.UpdatePostController {
	caseUpdate := usecases.NewUpdatePostUseCase(&mySQL)
	return controllers.NewUpdatePostController(caseUpdate)
}

func GetDeletePostController() *controllers.DeletePostController {
	caseDelete := usecases.NewDeletePostUseCase(&mySQL)
	return controllers.NewDeletePostController(caseDelete)
}
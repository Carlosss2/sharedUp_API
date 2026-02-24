package dependencies

import (
	"sharedup/app/src/helpers"
	usecases "sharedup/app/src/posts/application/useCases"
	"sharedup/app/src/posts/infraestructure"
	"sharedup/app/src/posts/infraestructure/adapters/realtime"
	"sharedup/app/src/posts/infraestructure/controllers"
)

var (
	mySQL infraestructure.MySQL
	hub *realtime.Hub
)

func InitPost() {

	db, err := helpers.ConectToMySQL()
	if err != nil {
		panic(err)
	}

	mySQL = *infraestructure.NewMySQL(db)

	// Inicializar Hub
	hub = realtime.NewHub()
	go hub.Run()
}

func GetCreatePostController() *controllers.CreatePostController {

	notifier := realtime.NewPostNotifier(hub)

	caseCreate := usecases.NewCreatePostUseCase(&mySQL, notifier)

	return controllers.NewCreatePostController(caseCreate)
}

func GetWebSocketHandler() *realtime.WebSocketHandler {
	return realtime.NewWebSocketHandler(hub)
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
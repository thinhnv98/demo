package routes

import (
	"database/sql"
	"demo/controllers"
	"demo/repositories"
	"demo/services"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Db     *sql.DB
	Server gin.Engine
}

func (_self Route) Register() {
	//Common
	userHandler := controllers.UserController{
		IUserService: services.UserService{
			IUserRepo: repositories.UserRepo{
				Db: _self.Db,
			},
		},
	}

	userRoutes := _self.Server.Group("/api")
	{
		userRoutes.GET("/ping", userHandler.Register)
		//userRoutes.POST()
	}
}

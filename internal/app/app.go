package app

import (
	"gaston.frassatti/aouth_manager/internal/client"
	"gaston.frassatti/aouth_manager/internal/database"
	"gaston.frassatti/aouth_manager/internal/handlers"
	"gaston.frassatti/aouth_manager/internal/repository"
	"gaston.frassatti/aouth_manager/internal/service"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Server *gin.Engine
}

func NewApplication() *Application {
	router := gin.Default()

	mysql := database.NewMySql()
	dbConnection := mysql.OpenConnection()
	dao := repository.NewGrantsDAO(dbConnection)
	client := client.NewSlingConfig().SlingClient()
	service := service.NewGrantsService(dao, client)
	handler := handlers.NewHandler(service)

	v1 := router.Group("oauthManager/v1")
	v1.POST("/grants", handler.ManageGrants)

	return &Application{
		Server: router,
	}
}

package app

import (
	"gaston.frassatti/aouth_manager/internal/client"
	"gaston.frassatti/aouth_manager/internal/database"
	"gaston.frassatti/aouth_manager/internal/handlers"
	"gaston.frassatti/aouth_manager/internal/service"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Server *gin.Engine
}

func NewApplication() *Application {
	urlBase := "https://94ff2c06e96fe23b7913b17145f3c1db.m.pipedream.net"
	config := database.Config{
		Driver:   "mysql",
		User:     "user",
		Pass:     "user",
		Database: "oauth_db?multiStatements=true",
		Port:     "@tcp(127.0.0.1:3306)/",
	}

	router := gin.Default()
	db := database.NewConnection(config)
	httpClient := client.NewHttpClient(urlBase)
	srv := service.NewGrantsService(db, httpClient)

	handler := handlers.NewHandler(srv)

	v1 := router.Group("oauthManager/v1")
	v1.POST("/grants", handler.GrantsHandler)

	return &Application{
		Server: router,
	}
}

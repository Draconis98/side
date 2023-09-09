package main

import (
	"log"
	"main/api"
	"main/database"
	"main/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	database.InitDBConnection()
	service.InitKubeClient()
	// TODO: the db will not close here if <C-c>.
	// Try to fix it.
	defer database.CloseDBConnection()
	defer log.Println("Close DB")

	r.GET("index", api.Index)
	r.GET("api/container", api.GetContainer)
	r.POST("api/container/new", api.CreateContainer)
	//r.POST("api/container/expand", api.ExpandContainer)
	r.POST("api/container/delete", api.DeleteContainer)

	log.Println("server stared.")

	r.Run(":8000")
}

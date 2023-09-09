package main

import (
	"log"
	"main/api"
	"main/database"
	"main/service"
	"main/utils"
	"net/http"

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

	r.LoadHTMLGlob("templates/*")
	r.GET("index", Index)
	r.GET("api/container", api.GetContainer)
	r.POST("api/container/new", api.CreateContainer)
	r.POST("api/container/expand", api.ExpandContainer)
	r.POST("api/container/delete", api.DeleteContainer)

	log.Println("server stared.")

	r.Run(":8000")
}

func Index(c *gin.Context) {
	var headerInfo utils.HeaderInfo
	c.BindHeader(&headerInfo)

	// Ensure the user exist in database.
	// So the user can create new container.
	if !database.CheckUserExists(headerInfo.Username) {
		log.Printf("Creating user %+v\n", headerInfo)
		database.InsertUser(headerInfo.Username)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{})
}

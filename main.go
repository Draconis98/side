package main

import (
	"github.com/gin-gonic/gin"
	 "log"
	"main/api"
)

func main() {
	r := gin.Default()

	r.GET("index", api.Index)
	//r.POST("api/container", api.GetContainer)
	r.POST("api/container/new", api.CreateContainer)
	//r.POST("api/container/expand", api.ExpandContainer)
	//r.POST("api/container/delete", api.DeleteContainer)

	log.Println("server stared.")

	r.Run(":8000")
}

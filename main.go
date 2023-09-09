package main

import (
	"github.com/gin-gonic/gin"
	_ "log"
	"main/api"
)

func main() {
	r := gin.Default()

	r.GET("index", api.Index)
	//r.POST("api/container", GetContainer)
	//r.POST("api/container/new", CreateContainer)
	//r.POST("api/container/expand", ExpandContainer)
	//r.POST("api/container/delete", DeleteContainer)

	r.Run(":8000")
}

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("index", Index)
	r.POST("api/container", GetContainer)
	r.POST("api/container/new", CreateContainer)
	r.POST("api/container/expand", ExpandContainer)
	r.POST("api/container/delete", DeleteContainer)

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

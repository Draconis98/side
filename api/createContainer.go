package api

import (
	"log"
	"github.com/gin-gonic/gin"
	"main/utils"
	_ "net/http"
)

func CreateContainer(c *gin.Context) {
	var creation utils.ContainerCreation
	var headerInfo utils.HeaderInfo
	c.BindJSON(&creation)
	c.BindHeader(&headerInfo)
	log.Printf("%+v\n", c.Request)
	log.Printf("%+v %+v\n", creation, headerInfo)

}

package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/utils"
)

func DeleteContainer(c *gin.Context) {
	var deletion utils.ContainerDeletion
	var headerInfo utils.HeaderInfo
	c.BindJSON(&deletion)
	c.BindHeader(&headerInfo)
	log.Printf("%+v\n", c.Request)
	log.Printf("%+v %+v\n", deletion, headerInfo)

}

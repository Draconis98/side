package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/database"
	"main/utils"
	"net/http"
	_ "net/http"
)

func GetContainer(c *gin.Context) {
	var headerInfo utils.HeaderInfo
	c.BindHeader(&headerInfo)
	log.Printf("Get container: %+v\n", headerInfo)

	// Get container list for current user.
	containerList := database.GetContainersByUser(headerInfo.Username)

	c.JSON(http.StatusOK, containerList)

}

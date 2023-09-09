package api

import (
	"log"
	"main/database"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetContainer(c *gin.Context) {
	var headerInfo utils.HeaderInfo
	c.BindHeader(&headerInfo)
	log.Printf("Get container: %+v\n", headerInfo)

	// Get container list for current user.
	containerList := database.GetContainersByUser(headerInfo.Username)
	if containerList == nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage{
			Message: "Username not exists, get container list failed.",
		})
	} else {
		c.JSON(http.StatusOK, containerList)
	}
}

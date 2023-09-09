package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/database"
	"main/service"
	"main/utils"
	"net/http"
)

func DeleteContainer(c *gin.Context) {
	var deletion utils.ContainerDeletion
	var headerInfo utils.HeaderInfo
	c.BindJSON(&deletion)
	c.BindHeader(&headerInfo)
	log.Printf("Delete container %+v %+v\n", deletion, headerInfo)

	// Delete from database first
	if flag := database.DeleteContainer(deletion.ContainerId); !flag {
		c.JSON(http.StatusInternalServerError, utils.ErrorMessage{Message: "Database error"})
		return
	}

	c.Status(http.StatusOK)

	// Then delete from k8s
	service.Delete(headerInfo.Username, deletion.ContainerId)
}

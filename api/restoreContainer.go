package api

import (
	"main/database"
	"main/service"
	"main/utils"

	"github.com/gin-gonic/gin"
)

func RestoreContainer(r *gin.Context) {
	var restore utils.ContainerRestore
	var headerInfo utils.HeaderInfo

	r.BindJSON(&restore)
	r.BindHeader(&headerInfo)

	name := restore.ContainerId
	core, mem := restore.Core, restore.Memory

	if err := service.RestoreDeployment(name, core, mem); err != nil {
		r.JSON(500, gin.H{
			"message": "Failed to restore the container"})
		return
	}

	service.CheckEndLoading(headerInfo.Username, name)
	database.UpdateContainerStatus(name, 1)

	r.JSON(200, gin.H{
		"containerId": name,
		"message":     "Restored the container"})
}

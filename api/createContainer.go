package api

import (
	"log"
	"main/database"
	"main/service"
	"main/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateContainer(c *gin.Context) {
	var creation utils.ContainerCreation
	var headerInfo utils.HeaderInfo
	c.BindJSON(&creation)
	c.BindHeader(&headerInfo)
	log.Printf("Create container: %+v %+v\n", creation, headerInfo)

	// Get resource limitation.
	coreLimit, memoryLimit := database.GetResourceLimitByUser(headerInfo.Username)
	if coreLimit == -1 || memoryLimit == -1 {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage{
			Message: "Username not exists, get resource limitation failed.",
		})
		return
	}

	// Get container list for current user.
	containerList := database.GetContainersByUser(headerInfo.Username)
	core, memory := creation.Core, creation.Memory
	for _, container := range containerList {
		core += container.Core
		memory += container.Memory
	}
	if core > coreLimit || memory > memoryLimit {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage{
			Message: "Exceed resource limitation, create container failed.",
		})
		return
	}

	createdAt := time.Now()
	createdAtStr := createdAt.Format("20060102150405")

	// Attention: the containerID should be the save as that in k8s.
	// Here is a misc, do not change it.
	containerId := strings.ToLower(creation.BaseImage) + "-" + createdAtStr + "-" + headerInfo.Username

	// This function will always success (maybe take a long time).
	service.Run(
		headerInfo.Username,
		creation.BaseImage,
		createdAtStr,
		creation.Core,
		creation.Memory,
	)
	// Write to database
	database.InsertContainer(
		containerId,
		headerInfo.Username,
		creation.BaseImage,
		1,
		creation.Core,
		creation.Memory,
		createdAt,
	)
	c.JSON(http.StatusOK, utils.Container{
		ContainerId: containerId,
		Core:        creation.Core,
		Memory:      creation.Memory,
		Status:      1,
		CreateAt:    createdAtStr,
	})
}

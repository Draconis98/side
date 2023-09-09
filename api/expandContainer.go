package api

import (
	"log"
	"main/database"
	"main/service"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExpandContainer(c *gin.Context) {
	var expansion utils.ContainerExpansion
	var headerInfo utils.HeaderInfo
	c.BindJSON(&expansion)
	c.BindHeader(&headerInfo)
	log.Printf("Expand container: %+v %+v\n", expansion, headerInfo)

	// Get resource limitation.
	coreLimit, memoryLimit := database.GetResourceLimitByUser(headerInfo.Username)
	if coreLimit == -1 || memoryLimit == -1 {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage{
			Message: "Username not exists, get resource limitation failed.",
		})
		return
	}

	coreOld, memoryOld := database.GetResourceInfoByContainerId(expansion.ContainerId)
	if coreOld == -1 || memoryOld == -1 {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage{
			Message: "Container not exists or get ContainerInfo failed.",
		})
		return
	}

	if coreOld > expansion.NewCore || memoryOld > expansion.NewMemory {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage{
			Message: "Requsting resources is invalid. Expanding failed.",
		})
		return
	}

	if coreOld == expansion.NewCore && memoryOld == expansion.NewMemory {
		c.Status(http.StatusOK)
		return
	}

	// Get container list for current user.
	containerList := database.GetContainersByUser(headerInfo.Username)
	if containerList == nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage{
			Message: "Username not exists, get container list failed.",
		})
	}

	core, memory := expansion.NewCore, expansion.NewMemory
	for _, container := range containerList {
		core += container.Core
		memory += container.Memory
	}
	core = core - coreOld
	memory = memory - memoryOld
	if core > coreLimit || memory > memoryLimit {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage{
			Message: "Exceed resource limitation, create container failed.",
		})
		return
	}

	// This function maybe take a long time.
	flag := service.ExpandRequirement(
		expansion.ContainerId,
		headerInfo.Username,
		coreOld,
		memoryOld,
		expansion.NewCore,
		expansion.NewMemory,
	)
	if !flag {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage{
			Message: "Expand container failed.",
		})
		return
	}
	// Write to database
	database.UpdateContainerInfo(
		expansion.ContainerId,
		expansion.NewCore,
		expansion.NewMemory,
	)
	c.JSON(http.StatusOK, utils.ContainerExpansion{
		ContainerId: expansion.ContainerId,
		NewCore:     expansion.NewCore,
		NewMemory:   expansion.NewMemory,
	})
}

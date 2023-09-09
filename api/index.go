package api

import (
	"main/database"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var headerInfo utils.HeaderInfo
	c.BindHeader(&headerInfo)

	// Ensure the user exist in database.
	// So the user can create new container.
	if !database.CheckUserExists(headerInfo.Username) {
		database.InsertUser(headerInfo.Username)
	}

	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, `OK`)
}

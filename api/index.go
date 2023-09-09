package api

import (
	"github.com/gin-gonic/gin"
	_ "net/http"
)

func Index(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, `OK`)
}


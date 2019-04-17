package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleHome(c *gin.Context) {
	response := gin.H{
		"msg": "Welcome Home!",
	}

	c.JSON(http.StatusOK, response)
}

func main() {
	r := gin.Default()
	r.GET("/", handleHome)
	r.Run()
}

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetSample(c *gin.Context) {
	params := c.Query("hoge")
	token := c.Query("token")
	c.JSON(200, gin.H{
		"message": "get",
		"params":  params,
		"token":   token,
	})
}

package controllers

import "github.com/gin-gonic/gin"

func PostTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "post",
	})
}

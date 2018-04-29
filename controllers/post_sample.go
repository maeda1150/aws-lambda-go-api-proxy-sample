package controllers

import "github.com/gin-gonic/gin"

type JSONParams struct {
	Hoge string `form:"hoge" binding:"required"`
}

func PostSample(c *gin.Context) {
	var json JSONParams
	c.Bind(&json)
	c.JSON(200, gin.H{
		"message": "post",
		"params":  json.Hoge,
	})
}

package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type JSONParams struct {
	Hoge  string `form:"hoge" binding:"required"`
	Token string `form:"token" binding:"required"`
}

func PostSample(c *gin.Context) {
	log.Println("----------------------")
	log.Println(c.Request.Header)
	log.Println("----------------------")
	log.Println(c.Request.Header["Content-Type"])
	log.Println("----------------------")
	var json JSONParams
	c.Bind(&json)
	c.JSON(200, gin.H{
		"message": "post",
		"params":  json.Hoge,
		"token":   json.Token,
	})
}

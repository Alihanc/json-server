package main

import (
	helper "json-server/helper"

	"github.com/gin-gonic/gin"
)

func main() {
	var router *gin.Engine = gin.Default()
	router.GET("/bunker", helper.Get)
	router.GET("/bunker/:id", helper.GetByID)
	router.POST("/post", helper.Post)
	router.DELETE("/delete/:id", helper.Delete)
	router.PATCH("/update/:id", helper.Update)
	router.Run("localhost:8080")
}

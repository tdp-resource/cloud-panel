package config

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/config")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/", list)
		rg.POST("/", create)
		rg.GET("/:key", fetch)
		rg.PATCH("/:key", update)
		rg.DELETE("/:key", delete)
	}

}

package main

import (
	"code.gopub.tech/gbt/handler"
	"github.com/gin-gonic/gin"
)

func register(r *gin.Engine) {
	api := r.Group("/api")
	api.GET("/ping", Wrap(handler.Ping))
}

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/ffxiv-todo/pkg/api"
)

func Install(r gin.IRouter) {
	r.POST("/calc", api.Calculate)
}

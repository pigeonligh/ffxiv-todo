package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func response(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}

func failed(err error) gin.H {
	return gin.H{
		"errcode": 1,
		"errmsg":  fmt.Sprint(err),
	}
}

func success(obj gin.H) gin.H {
	obj["errcode"] = 0
	obj["errmsg"] = ""
	return obj
}

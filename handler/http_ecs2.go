package handler

import (
	"github.com/gin-gonic/gin"
)

func TestSimpleApiWithPort(path string, data []byte) *gin.Engine {
	r := gin.Default()
	r.GET(path, func(c *gin.Context) {
		c.Data(200, "json", data)
	})
	return r
}

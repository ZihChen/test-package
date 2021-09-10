package router

import (
	"github.com/gin-gonic/gin"
	"test-package-02/app/handler/stock"
)

func RouteProvider(r *gin.Engine) {
	stockh := stock.NewHandler()

	v1 := r.Group("v1")
	{
		v1.POST("/", stockh.Create)
		v1.GET("/", stockh.Read)
	}
}
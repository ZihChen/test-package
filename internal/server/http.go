package server

import (
	"github.com/gin-gonic/gin"
	"test-package-02/router"
)

func Run() {
	r := gin.New()

	router.RouteProvider(r)

	r.Run(":8080")
}
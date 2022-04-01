package rest

import (
	"github.com/gin-gonic/gin"
	"log"
)

func server() {
	r := gin.Default()
	rg := r.Group("/api")
	rg.GET("/user/:id")
	rg.GET("/user/:name")
	rg.POST("/user")
	rg.PUT("/user")
	rg.DELETE("/user/:id")
	rg.POST("/user/auth")
	log.Panic(r.Run(":9000"))
}

package router

import (
	"github.com/gin-gonic/gin"
)



func RegisterUserRoutes(e *gin.Engine) {
	routes := e.Group("/api")
	//routes.Use(authMiddleWare.AuthUser())
	routes.GET("/users", func(ctx *gin.Context) {})
	
}

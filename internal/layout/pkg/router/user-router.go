package router

import (
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(e *gin.Engine, module *user.Module) {
	routes := e.Group("/api")
	//routes := gin.Default()
	//routes.Use(authMiddleWare.AuthUser())
	routes.GET("/users", module.Handler.FindAllUser)
	routes.POST("/users", module.Handler.CreateUser)
	routes.GET("/users/:id", module.Handler.FindUser)
	routes.DELETE("/users/:id", module.Handler.DeleteUser)

}

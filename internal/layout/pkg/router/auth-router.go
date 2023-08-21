package router

import (
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/auth"
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/pkg/constant"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(e *gin.Engine, module *auth.Module) {
	routes := e.Group(constant.ApiPattern + constant.V1 + constant.AuthRouteName)
	routes.POST(constant.RootPattern, module.Handler.Login)

}

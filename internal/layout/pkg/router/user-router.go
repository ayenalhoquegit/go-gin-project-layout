package router

import (
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user"
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/pkg/constant"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(e *gin.Engine, module *user.Module) {
	routes := e.Group(constant.ApiPattern+constant.V1+constant.UserRouteName)
	//routes := gin.Default()
	//routes.Use(authMiddleWare.AuthUser())
	routes.GET(constant.RootPattern, module.Handler.FindAllUser)
	routes.POST(constant.RootPattern, module.Handler.CreateUser)
	routes.GET(constant.RootPattern+":id", module.Handler.FindUser)
	routes.PATCH(constant.RootPattern+":id", module.Handler.UpdateUser)
	routes.DELETE(constant.RootPattern+":id", module.Handler.DeleteUser)

}

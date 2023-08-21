package layout

import (
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/auth"
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user"
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/pkg"
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/pkg/db"
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/pkg/router"
)

type App struct {
	DBClient   *db.DbClient
	gin        *pkg.Gin
	UserModule *user.Module
	AuthModule *auth.Module
}

func NewApp() *App {
	a := new(App)
	a.initComponent()
	return a
}

func (a *App) initDB() {
	a.DBClient = db.GetDbInstance()
}

func (a *App) initModules() {
	a.UserModule = user.NewModule(a.DBClient.DB)
	a.AuthModule = auth.NewModule(a.UserModule.Service)
}

func (a *App) initModuleRouters() {
	router.RegisterUserRoutes(a.gin.Engine, a.UserModule)
	router.RegisterAuthRoutes(a.gin.Engine, a.AuthModule)
}

func (a *App) initComponent() {
	a.initDB()
	a.gin = pkg.NewGin()
	a.initModules()
	a.initModuleRouters()

}

// Run app
func (a *App) Run() {
	a.gin.Engine.Run(":8080")
}

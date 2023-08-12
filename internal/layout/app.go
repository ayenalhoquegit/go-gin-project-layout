package layout

import (
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user"
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/pkg"
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/pkg/db"
)

type App struct{
	DBClient *db.DbClient
	gin *pkg.Gin
	UserModule *user.Module

}
func NewApp() * App{
	a:=new(App)
	a.initComponent()
	return a
}

func (a *App) initDB(){
	a.DBClient = db.GetDbInstance()
}

func (a *App) initModules(){
	a.UserModule = user.NewModule(a.DBClient)
}

func (a *App) initComponent(){
	a.initDB()
	a.gin = pkg.NewGin()	
	a.initModules()
	
}

// Run app
func (a *App) Run() {
	a.gin.Engine.Run(":8080")
}



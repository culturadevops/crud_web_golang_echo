package main

import (
	"github.com/culturadevops/echoweb/libs"
	"github.com/culturadevops/echoweb/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	dbConfig := libs.Configure("./configs", "mysql")
	libs.DB = dbConfig.InitMysqlDB()
}
func mainMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
func mainRoutes(e *echo.Echo) {
	routes.Notas(e)

}
func main() {
	e := echo.New()
	mainMiddlewares(e)
	mainRoutes(e)
	e.Logger.Fatal(e.Start(":8081"))
}

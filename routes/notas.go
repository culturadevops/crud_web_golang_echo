package routes

import (
	"github.com/culturadevops/echoweb/controllers"
	"github.com/labstack/echo"
)

func Notas(e *echo.Echo) {

	/*middlewares*/
	e.GET("/", controllers.Notas.List)
	e.POST("/", controllers.Notas.Add)
	e.DELETE("/:id", controllers.Notas.Del)
	e.PUT("/:id", controllers.Notas.Update)
}

package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/culturadevops/echoweb/models"
	"github.com/labstack/echo"
)

var Notas *notasController

type notasController struct {
}

type RequestNotas struct {
	Title       string `json:"title" form:"title" query:"title"`
	Description string `json:"description" form:"description" query:"description"`
}

func (t *notasController) Add(c echo.Context) error {
	u := new(RequestNotas)
	if err := c.Bind(u); err != nil {
		return err
	}
	models.Notas.Add(u.Title, u.Description)
	return c.JSON(http.StatusOK, "Registro Creado")
}
func (t *notasController) List(c echo.Context) error {
	list := models.Notas.List()
	return c.JSON(http.StatusOK, list)
}
func (t *notasController) Show(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	use, _ := models.Notas.Info(uint(id))
	return c.JSON(http.StatusOK, use)
}

func (t *notasController) Del(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	use := models.Notas.Del(int(id))
	fmt.Println(use)
	return c.JSON(http.StatusOK, "Registro eliminado")
}

func (t *notasController) Update(c echo.Context) error {

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	u := new(RequestNotas)
	if err := c.Bind(u); err != nil {
		return err
	}
	models.Notas.Update(int(id), u.Title, u.Description)

	return c.JSON(http.StatusOK, "Registro actualizado")
}

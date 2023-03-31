package models

import (
	"errors"
	"log"

	"github.com/culturadevops/echoweb/libs"
	"github.com/jinzhu/gorm"
)

var Notas *notas

type notas struct {
	gorm.Model
	Title       string
	Description string
}

func (this *notas) List() []notas {
	var data = []notas{}
	err := libs.DB.Find(&data).Error
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func (this *notas) Info(id uint) (notas, error) {
	var data notas

	if libs.DB.Where("id  = ? ", id).Find(&data).RecordNotFound() {
		return notas{}, errors.New("No existe registro con es id")
	}

	return data, nil
}

func (this *notas) Add(title string, description string) (notas, error) {
	var data notas
	data.Title = title
	data.Description = description
	if err := libs.DB.Create(&data).Error; err != nil {
		return notas{}, err
	} else {

		return data, nil
	}

}

func (this *notas) Update(id int, title string, description string) (notas, error) {
	var data notas

	if libs.DB.Where("id = ? ", id).First(&data).RecordNotFound() {
		return notas{}, errors.New("No hay registro con ese id")
	}
	data.Title = title
	data.Description = description
	if err := libs.DB.Save(&data).Error; err != nil {
		return notas{}, errors.New("no se pudo actualizar")
	}
	return data, nil

}

func (this *notas) Del(id int) error {
	var data notas
	if err := libs.DB.Where("id = ?", id).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}

package repository

import (
	"golang-micro/configs"
	"golang-micro/entity"
)

func GetAllLaptop() ([]entity.Laptop, error) {
	var laptop []entity.Laptop
	err := configs.Dbcon.Find(&laptop).Error
	return laptop, err
}

func GetOneLaptopById(id int) (entity.Laptop, error) {
	var laptop entity.Laptop
	err := configs.Dbcon.Where("id = ?", id).First(&laptop).Error
	return laptop, err
}

func CreateLaptop(laptop entity.Laptop) (entity.Laptop, error) {
	err := configs.Dbcon.Create(&laptop).Error
	return laptop, err
}

func UpdateOneLaptopById(laptop entity.Laptop) (entity.Laptop, error) {
	err := configs.Dbcon.Save(&laptop).Error
	return laptop, err
}

func DeleteOneLaptopById(id int) {
	var laptop entity.Laptop
	configs.Dbcon.Where("id = ?", id).Delete(&laptop)
	return
}

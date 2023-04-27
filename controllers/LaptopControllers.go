package controllers

import (
	"golang-micro/dto"
	"golang-micro/entity"
	"golang-micro/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	data, err := repository.GetAllLaptop()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read get data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func Show(c *gin.Context) {
	id := c.Param("id")
	converIdToInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed convert data"})
		return
	}

	data, err := repository.GetOneLaptopById(converIdToInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read get data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
func Create(c *gin.Context) {
	var req dto.LaptopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	obj := entity.Laptop{
		Nama:  req.Nama,
		Harga: req.Harga,
		Merk:  req.Merk,
		Os:    req.Os,
	}

	data, err := repository.CreateLaptop(obj)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data, "message": "data created"})
}

func Update(c *gin.Context) {
	var req dto.LaptopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	converIdToInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed convert data"})
		return
	}
	tempdata, err := repository.GetOneLaptopById(converIdToInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read get data"})
		return
	}
	if req.Nama != "" {
		tempdata.Nama = req.Nama
	}
	if req.Harga != 0 {
		tempdata.Harga = req.Harga
	}
	if req.Merk != "" {
		tempdata.Merk = req.Merk
	}
	if req.Os != "" {
		tempdata.Os = req.Os
	}
	data, err := repository.UpdateOneLaptopById(tempdata)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read get data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	converIdToInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed convert data"})
		return
	}
	repository.DeleteOneLaptopById(converIdToInt)
	c.JSON(http.StatusOK, gin.H{"data": "Laptop Deleted"})
}

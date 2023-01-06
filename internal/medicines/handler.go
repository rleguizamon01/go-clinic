package medicines

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetMedicinesHandler(c *gin.Context) {
	medicines, err := GetMedicines()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, medicines)
}

func GetMedicineHandler(c *gin.Context) {
	medicine, err := GetMedicine(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, medicine)
}

func CreateMedicineHandler(c *gin.Context) {
	var medicine Medicine

	if err := c.Bind(&medicine); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	newMedicine, err := CreateMedicine(medicine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, newMedicine)
}

func UpdateMedicineHandler(c *gin.Context) {
	var medicine Medicine

	if err := c.Bind(&medicine); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	medicine.ID = uint(id)
	medicine, err = UpdateMedicine(medicine)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, medicine)
}

func DeleteMedicineHandler(c *gin.Context) {
	err := DeleteMedicine(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

package controllers

import (
	"APIs/src/mpu6050/domain/entities"
	"APIs/src/mpu6050/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Mpu6050Controller struct {
    mpu6050Service *application.Mpu6050Service
}

func NewMpu6050Controller(mpu6050Service *application.Mpu6050Service) *Mpu6050Controller {
    return &Mpu6050Controller{mpu6050Service: mpu6050Service}
}

func (c *Mpu6050Controller) SaveMPU6050Data(ctx *gin.Context) {
    var data entities.MPU6050
    if err := ctx.ShouldBindJSON(&data); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.mpu6050Service.SaveMPU6050Data(&data); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Data saved successfully"})
}
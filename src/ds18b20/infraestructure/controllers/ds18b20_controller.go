package controllers

import (
    "net/http"
    "APIs/src/ds18b20/application"
    "APIs/src/ds18b20/domain/entities"
    "github.com/gin-gonic/gin"
)

type Ds18b20Controller struct {
    ds18b20Service *application.Ds18b20Service
}

func NewDS18B20Controller(ds18b20Service *application.Ds18b20Service) *Ds18b20Controller {
    return &Ds18b20Controller{ds18b20Service: ds18b20Service}
}

func (c *Ds18b20Controller) SaveDS18B20Data(ctx *gin.Context) {
    var data ds_entities.DS18B20
    if err := ctx.ShouldBindJSON(&data); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.ds18b20Service.SaveDS18B20Data(&data); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Data saved successfully"})
}

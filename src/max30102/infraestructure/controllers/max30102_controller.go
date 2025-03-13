package controllers

import (
    "net/http"
    "APIs/src/max30102/application"
    "APIs/src/max30102/domain/entities"
    "github.com/gin-gonic/gin"
)

type Max30102Controller struct {
    max30102Service *application.Max30102Service
}

func NewMax30102Controller(max30102Service *application.Max30102Service) *Max30102Controller {
    return &Max30102Controller{max30102Service: max30102Service}
}

func (c *Max30102Controller) SaveMax30102Data(ctx *gin.Context) {
    var data entities.Max30102
    if err := ctx.ShouldBindJSON(&data); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.max30102Service.SaveMax30102Data(&data); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Data saved"})
}

func (c *Max30102Controller) GetMax30102Data(ctx *gin.Context) {
    data, err := c.max30102Service.GetMax30102Data()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"data": data})
}
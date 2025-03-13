package controllers

import (
    "net/http"
    "APIs/src/max30102/application"
    "github.com/gin-gonic/gin"
)

type RespuestaController struct {
    respuestaService *application.RespuestaService
}

func NewRespuestaController(respuestaService *application.RespuestaService) *RespuestaController {
    return &RespuestaController{respuestaService: respuestaService}
}

func (c *RespuestaController) SendMessage(ctx *gin.Context) {
    var request struct {
        Queue   string `json:"queue"`
        Message string `json:"message"`
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.respuestaService.SendMessage(request.Queue, request.Message); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}
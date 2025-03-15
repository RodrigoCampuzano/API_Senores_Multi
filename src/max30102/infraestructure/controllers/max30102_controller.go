package controllers

import (
	"APIs/src/max30102/application"
	"APIs/src/max30102/domain/entities"
	"APIs/src/max30102/infraestructure/broker"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateMax30102Controller struct {
    createUseCase *application.CreateMax30102UseCase
    publisher     *broker.RabbitMQPublisher
}

func NewMax30102Controller(max30102Service *application.CreateMax30102UseCase, publisher *broker.RabbitMQPublisher) *CreateMax30102Controller {
    return &CreateMax30102Controller{createUseCase: max30102Service}
}

func (c *CreateMax30102Controller) SaveMax30102Data(ctx *gin.Context) {
    var data *entities.Max30102
    if err := ctx.ShouldBindJSON(&data); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.createUseCase.Run(data); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    message := map[string]interface{}{
		"bpm":   data.BPM,
		"spo2": data.SpO2,
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Println("Error al serializar el mensaje:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar el mensaje"})
		return
	}
	err = c.publisher.Publish(messageBytes)
	if err != nil {
		log.Println("Error publicando mensaje en RabbitMQ:", err)
	}
    ctx.JSON(http.StatusOK, gin.H{"message": "Data saved"})
}


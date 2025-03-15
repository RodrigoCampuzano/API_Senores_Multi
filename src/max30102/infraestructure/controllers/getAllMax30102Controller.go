package controllers

import (
	"APIs/src/max30102/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllMax30102Controller struct {
	getAllUseCase *application.GetAll30102UseCase
}

func GetAll30102Controller(max30102Service *application.GetAll30102UseCase) *GetAllMax30102Controller {
	return &GetAllMax30102Controller{getAllUseCase: max30102Service}
}

func (c *GetAllMax30102Controller) GetMax30102Data(ctx *gin.Context) {
	data, err := c.getAllUseCase.GetMax30102Data()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

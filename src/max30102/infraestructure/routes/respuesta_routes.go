package routes

import (
    "APIs/src/max30102/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRespuestaRoutes(r *gin.Engine, respuestaController *controllers.RespuestaController) {
    respuestaGroup := r.Group("/respuesta")
    {
        respuestaGroup.POST("/send", respuestaController.SendMessage)
    }
}	
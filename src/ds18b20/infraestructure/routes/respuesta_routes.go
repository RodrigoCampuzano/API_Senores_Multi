package routes

import (
    "APIs/src/ds18b20/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRespuestaRoutes(r *gin.Engine, respuestaController *controllers.RespuestaController) {
    respuestaGroup := r.Group("/respuesta")
    {
        respuestaGroup.POST("/send", respuestaController.SendMessage)
    }
}	
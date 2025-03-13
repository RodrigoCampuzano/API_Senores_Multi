package routes

import (
    "APIs/src/ds18b20/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

func DS18B20Routes(r *gin.Engine, ds18b20Controller *controllers.Ds18b20Controller) {
    healthGroup := r.Group("/ds18b20")
    {
        healthGroup.POST("/", ds18b20Controller.SaveDS18B20Data)
    }
}
package routes

import (
    "APIs/src/mpu6050/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

func MPU6050Routes(r *gin.Engine, Mpu6050Controller *controllers.Mpu6050Controller) {
    healthGroup := r.Group("/mpu6050")
    {
        healthGroup.POST("/", Mpu6050Controller.SaveMPU6050Data)
    }
}
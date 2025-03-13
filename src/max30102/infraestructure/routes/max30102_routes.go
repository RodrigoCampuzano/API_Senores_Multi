package routes

import (
    "APIs/src/max30102/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

func Max30102Routes(r *gin.Engine, max30102Controller *controllers.Max30102Controller) {
    max30102Group := r.Group("/max30102")
    {
        max30102Group.POST("/", max30102Controller.SaveMax30102Data)
    }
}
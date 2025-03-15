package infraestructure

import (
	"github.com/gin-gonic/gin"
)

func Max30102Routes(r *gin.Engine, max30102Service *ProductController) {
	max30102Group := r.Group("/max30102")
	{
		max30102Group.POST("/", max30102Service.CreateProduct)
		max30102Group.GET("/", max30102Service.GetAllProducts)
	}
}

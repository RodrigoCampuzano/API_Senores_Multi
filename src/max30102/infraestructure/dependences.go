package infraestructure

import (
	rabbitmq "APIs/src/core/rabbitMQ"
	"APIs/src/max30102/application"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	ps := NewMySQL()
	createProduct := application.NewMax30102UseCase(ps)
	getAllProducts := application.NewGetAllMax30102UseCase(ps)
	publisher := rabbitmq.RabbitConnection()
	productController := NewProductController(createProduct, getAllProducts, publisher)
	Max30102Routes(r, productController)
}

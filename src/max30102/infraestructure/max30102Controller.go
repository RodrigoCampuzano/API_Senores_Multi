package infraestructure

import (
	"APIs/src/max30102/application"
	"APIs/src/max30102/infraestructure/broker"
	"APIs/src/max30102/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	createProductUseCase   *controllers.CreateMax30102Controller
	viewAllProductsUseCase *controllers.GetAllMax30102Controller
}

func NewProductController(
	createUseCase *application.CreateMax30102UseCase,
	viewUseCase *application.GetAll30102UseCase,
	publisher *broker.RabbitMQPublisher, // 🔹 Se agrega el publisher de RabbitMQ
) *ProductController {
	// 🔹 Se pasa el publisher al handler de creación
	createHandler := controllers.NewMax30102Controller(createUseCase, publisher)
	viewHandler := controllers.GetAll30102Controller(viewUseCase)

	return &ProductController{
		createProductUseCase:   createHandler,
		viewAllProductsUseCase: viewHandler,
	}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	pc.createProductUseCase.SaveMax30102Data(c)
}

func (pc *ProductController) GetAllProducts(c *gin.Context) {
	pc.viewAllProductsUseCase.GetMax30102Data(c)
}

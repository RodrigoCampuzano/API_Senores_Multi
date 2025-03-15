package infraestructure

import (
	"APIs/src/max30102/application"
	"APIs/src/max30102/infraestructure/broker"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func Init(r *gin.Engine) {
	ps := NewMySQL()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}
	createProduct := application.NewMax30102UseCase(ps)
	getAllProducts := application.NewGetAllMax30102UseCase(ps)
	host := os.Getenv("BROKER_HOST")
	user := os.Getenv("BROKER_USER")
	pass := os.Getenv("BROKER_PASS")
	connURL := fmt.Sprintf("amqp://%s:%s@%s/", user, pass, host)
	conn, err := amqp.Dial(connURL)
	if err != nil {
		log.Fatal("Error al conectar con RabbitMQ", err)
	}

	publisher, err := broker.NewRabbitMQPublisher(conn, "Q1")
	if err != nil {
		log.Fatal("Error al crear el publicador de RabbitMQ", err)
	}

	productController := NewProductController(createProduct, getAllProducts, publisher)

	Max30102Routes(r, productController)
}

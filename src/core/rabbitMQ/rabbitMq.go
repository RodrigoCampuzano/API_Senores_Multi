package rabbitmq

import (
	"APIs/src/max30102/infraestructure/broker"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func RabbitConnection () *broker.RabbitMQPublisher {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}
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
	return publisher
}
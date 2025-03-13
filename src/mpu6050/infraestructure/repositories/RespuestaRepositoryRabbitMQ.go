package repositories

import (
    "log"
    "github.com/streadway/amqp"
)

type RespuestaRepositoryRabbitMQ struct {
    conn *amqp.Connection
}

func NewRespuestaRepositoryRabbitMQ(url string) (*RespuestaRepositoryRabbitMQ, error) {
    conn, err := amqp.Dial(url)
    if err != nil {
        return nil, err
    }
    return &RespuestaRepositoryRabbitMQ{conn: conn}, nil
}

func (r *RespuestaRepositoryRabbitMQ) SendMessage(queue string, message string) error {
    ch, err := r.conn.Channel()
    if err != nil {
        return err
    }
    defer ch.Close()

    q, err := ch.QueueDeclare(
        queue, // name
        false, // durable
        false, // delete when unused
        false, // exclusive
        false, // no-wait
        nil,   // arguments
    )
    if err != nil {
        return err
    }

    err = ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(message),
        })
    if err != nil {
        return err
    }

    log.Printf(" [x] Sent %s", message)
    return nil
}
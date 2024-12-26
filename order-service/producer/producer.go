package producer

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var count = 0
var conn *amqp.Connection

func PublishOrderMessage(orderID string) {
	for {
		connectionToRMQ, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err != nil {
			failOnError(err, "Failed to connect to RabbitMQ")
			count++
		} else {
			conn = connectionToRMQ
			break
		}

		if count > 5 {
			panic("cant to connect to rmq")
		}

		time.Sleep(3 * time.Second)
		log.Println("Backing off ...")
		continue
	}
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"orders", // имя обменника
		"fanout", // тип обменника
		true,     // устойчивый
		false,    // автоудаление
		false,    // внутренний
		false,    // без ожидания
		nil,      // аргументы
	)
	failOnError(err, "Failed to declare an exchange")

	body := "New Order ID: " + orderID
	err = ch.Publish(
		"orders", // имя обменника
		"",       // ключ маршрутизации
		false,    // обязательный
		false,    // немедленный
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

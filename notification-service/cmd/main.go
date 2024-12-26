package main

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// back off
	time.Sleep(10 * time.Second)
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Декларация очереди
	q, err := ch.QueueDeclare(
		"notifications", // имя очереди
		true,            // устойчивая
		false,           // автоудаление
		false,           // эксклюзивная
		false,           // без ожидания
		nil,             // аргументы
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Привязка к обменнику
	err = ch.QueueBind(
		q.Name,   // имя очереди
		"",       // ключ маршрутизации
		"orders", // имя обменника
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // имя очереди
		"",     // имя потребителя
		true,   // автоподтверждение
		false,  // эксклюзивный
		false,  // не локальный
		false,  // без ожидания
		nil,    // аргументы
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}
	log.Printf(" [*] Waiting for messages in %s. To exit press CTRL+C", q.Name)
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] Received: %s", d.Body)
			log.Printf(" [✓] Processed notification!")
		}
	}()

	<-forever
}

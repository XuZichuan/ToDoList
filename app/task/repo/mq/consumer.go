package mq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeMessage(_ context.Context, queueName string) (msgs <-chan amqp.Delivery, err error) {
	ch, err := RabbitMq.Channel()
	if err != nil {
		return msgs, err
	}
	q, _ := ch.QueueDeclare(queueName, true, false, false, false, nil)
	err = ch.Qos(1, 0, false)
	return ch.Consume(q.Name, "", false, false, false, false, nil)

}

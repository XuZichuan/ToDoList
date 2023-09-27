package mq

import (
	"ToDoList/config"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitMq *amqp.Connection

func InitRabbitMQ() {
	//"amqp://guest:guest@localhost:5672/"
	connString := fmt.Sprintf("%s://%s:%s@%s:%s/", config.RabbitMQ, config.RabbitMQUser, config.RabbitMQPassWord, config.RabbitMQHost, config.RabbitMQPort)
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(err)
	}
	RabbitMq = conn
}

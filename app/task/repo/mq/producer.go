package mq

import (
	"ToDoList/consts"
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SendMessage2MQ(ctx context.Context, body []byte) error {
	//拿到消息队列实例
	ch, err := RabbitMq.Channel()
	if err != nil {
		return err
	}

	queue, err := ch.QueueDeclare(consts.RabbitMqTaskQueue, true, false, false, false, nil)
	if err != nil {
		return err
	}
	// 发送消息
	fmt.Println(queue.Name)
	err = ch.PublishWithContext(ctx, "", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		return err
	}
	fmt.Println("发送成功")
	return nil
}

package task

import (
	"ToDoList/app/task/repo/mq"
	"ToDoList/app/task/service"
	"ToDoList/consts"
	log "ToDoList/logger"
	"ToDoList/proto/pb"
	"context"
	"encoding/json"
)

//另外构建脚本在初始化的时候就开始运行。
type SycnTask struct{}

func (s *SycnTask) RunTaskCreate(ctx context.Context) error {
	var forever chan struct{}
	msgs, err := mq.ConsumeMessage(ctx, consts.RabbitMqTaskQueue)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			reqRabbitMQ := &pb.TaskRequest{}
			err = json.Unmarshal(msg.Body, reqRabbitMQ)
			if err != nil {
				return
			}
			//循环取出消息,落库
			err = service.TaskMQ2MySQL(ctx, reqRabbitMQ)
			if err != nil {
				return
			}
			msg.Ack(false)

		}
	}()

	log.LogrusObj.Infoln(err)
	<-forever
	return nil
}

package script

import (
	"ToDoList/app/task/repo/mq/task"
	log "ToDoList/logger"
	"context"
)

func RunTaskSycn(ctx context.Context) {
	sycnTask := &task.SycnTask{}
	log.LogrusObj.Info("script:RunTaskSycn is running")
	err := sycnTask.RunTaskCreate(ctx)
	if err != nil {
		log.LogrusObj.Infof("RunTaskCreate:%s", err)
	}
}

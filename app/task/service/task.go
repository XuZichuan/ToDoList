package service

import (
	"ToDoList/app/task/repo/db/dao"
	"ToDoList/app/task/repo/db/model"
	"ToDoList/app/task/repo/mq"
	"ToDoList/consts"
	log "ToDoList/logger"
	"ToDoList/proto/pb"
	"context"
	"encoding/json"
	"sync"
)

var TaskSvcIns *TaskSvc
var TaskSvcOnce sync.Once

type TaskSvc struct {
}

func GetTaskSvc() *TaskSvc {
	TaskSvcOnce.Do(func() {
		TaskSvcIns = &TaskSvc{}
	})
	return TaskSvcIns
}

// CreateTask  mirco的规范：CreateTask(ctx context.Context, request *pb.TaskRequest, response *pb.TaskDetailResponse) error
func (t *TaskSvc) CreateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) error {
	//整点花活用个MQ
	resp.Code = consts.SUCCESS
	body, _ := json.Marshal(req)
	err := mq.SendMessage2MQ(ctx, body)
	if err != nil {
		resp.Code = consts.ERROR
		return err
	}
	return nil
}

func TaskMQ2MySQL(ctx context.Context, req *pb.TaskRequest) error {
	task := model.Task{
		Uid:       uint(req.Uid),
		Title:     req.Title,
		Status:    int(req.Status),
		Content:   req.Content,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	return dao.NewTaskDao(ctx).CreateTask(ctx, &task)
}

func (t *TaskSvc) GetTasksList(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskListResponse) error {
	resp.Code = consts.SUCCESS
	if req.Limit == 0 {
		req.Limit = 10 //初始化为10个防止sql出问题
	}
	taskList, count, err := dao.NewTaskDao(ctx).GetTasksList(ctx, req.Uid, req.Start, req.Limit)
	if err != nil {
		resp.Code = consts.ERROR
		return err
	}
	resp.TaskList = toPBTaskList(taskList)
	resp.Count = uint32(count)
	return nil
}

func toPBTaskList(taskList []*model.Task) []*pb.TaskModel {
	resp := make([]*pb.TaskModel, 0)
	for _, task := range taskList {
		resp = append(resp, toPBTask(task))
	}
	return resp
}

func toPBTask(task *model.Task) *pb.TaskModel {
	return &pb.TaskModel{
		Id:         uint64(task.ID),
		Uid:        uint64(task.Uid),
		Title:      task.Title,
		Content:    task.Content,
		StartTime:  task.StartTime,
		EndTime:    task.EndTime,
		Status:     int64(task.Status),
		CreateTime: task.CreatedAt.UnixMilli(),
		UpdateTime: task.UpdatedAt.UnixMilli(),
	}
}

// GetTask
func (t *TaskSvc) GetTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) error {
	resp.Code = consts.SUCCESS
	task, err := dao.NewTaskDao(ctx).GetTaskByUserIdAndTaskId(ctx, req.Uid, req.Id)
	if task.ID == 0 || err != nil {
		resp.Code = consts.ERROR
		return err
	}
	resp.TaskDetail = toPBTask(task)
	return nil
}

func (t *TaskSvc) UpdateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) error {
	// 查找该用户的这条信息
	resp.Code = consts.SUCCESS
	taskData, err := dao.NewTaskDao(ctx).UpdateTask(ctx, req.Uid, req.Id, req.Title, req.Content, req.Status)
	if err != nil {
		resp.Code = consts.ERROR
		log.LogrusObj.Error("UpdateTask err:%v", err)
		return err
	}
	resp.TaskDetail = toPBTask(taskData)
	return nil
}

func (t *TaskSvc) DeleteTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) error {
	resp.Code = consts.SUCCESS
	err := dao.NewTaskDao(ctx).DeleteTaskByUserIdAndTaskId(ctx, req.Uid, req.Id)
	if err != nil {
		resp.Code = consts.ERROR
		return err
	}
	return nil
}

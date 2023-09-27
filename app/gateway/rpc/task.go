package rpc

import (
	"ToDoList/consts"
	"ToDoList/proto/pb"
	"context"
)

func TaskCreate(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	r, err := TaskService.CreateTask(ctx, req)
	if err != nil {
		return
	}
	if r.Code != consts.SUCCESS {
		return
	}

	return r, nil
}

func TaskUpdate(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	r, err := TaskService.UpdateTask(ctx, req)
	if err != nil {
		return
	}
	if r.Code != consts.SUCCESS {
		return
	}

	return r, nil
}

func TaskDelete(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	r, err := TaskService.DeleteTask(ctx, req)
	if err != nil {
		return
	}
	if r.Code != consts.SUCCESS {
		return
	}

	return r, nil
}

func TaskList(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskListResponse, err error) {
	r, err := TaskService.GetTasksList(ctx, req)
	if err != nil {
		return
	}
	if r.Code != consts.SUCCESS {
		return
	}

	return r, nil
}

func TaskGet(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	r, err := TaskService.GetTask(ctx, req)
	if err != nil {
		return
	}
	if r.Code != consts.SUCCESS {
		return
	}

	return r, nil
}

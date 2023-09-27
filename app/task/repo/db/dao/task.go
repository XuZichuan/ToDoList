package dao

import (
	"context"
	"gorm.io/gorm"

	"ToDoList/app/task/repo/db/model"
)

type TaskDao struct {
	*gorm.DB
}

//在dao中拿到db的实例 然后操作数据库

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{
		NewDBClient(ctx),
	}
}

// 数据交互层
func (t *TaskDao) CreateTask(ctx context.Context, task *model.Task) error {
	db := t.WithContext(ctx).Model(&model.Task{}).Create(&task)
	return db.Error
}

//直接获取全部的
func (t *TaskDao) GetTasksList(ctx context.Context, uid uint64, start, limit uint32) ([]*model.Task, int64, error) {
	result := make([]*model.Task, 0, 0)
	count := int64(0)
	//获取详细数据
	db := t.WithContext(ctx).Model(&model.Task{})
	db.Where("uid = ?", uid)
	//获取翻页总数
	//db.Count(&count)
	//获取翻页数据
	db.Offset(int(start)).Limit(int(limit))
	db.Scan(&result)

	return result, count, db.Error
}

func (t *TaskDao) GetTaskByUserIdAndTaskId(ctx context.Context, uid, id uint64) (*model.Task, error) {
	result := &model.Task{}
	db := t.WithContext(ctx).Model(&model.Task{}).Where("uid = ?", uid)
	db.Where("id = ?", id)
	db.Scan(result)
	return result, db.Error
}

func (t *TaskDao) DeleteTaskByUserIdAndTaskId(ctx context.Context, uid, id uint64) error {
	return t.WithContext(ctx).Model(&model.Task{}).
		Where("id =? AND uid=?", id, uid).
		Delete(&model.Task{}).Error
}

func (t *TaskDao) UpdateTask(ctx context.Context, uid, id uint64, title, content string, status int64) (*model.Task, error) {
	result := &model.Task{}
	db := t.WithContext(ctx).Model(&model.Task{}).
		Where("id = ? AND uid = ?", id, uid).
		First(result)
	if db.Error != nil {
		return nil, db.Error
	}
	result.Title = title
	result.Status = int(status)
	result.Content = content
	t.Save(result)
	return result, db.Error
}

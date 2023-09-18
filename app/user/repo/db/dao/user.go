package dao

import (
	"ToDoList/app/user/repo/db/model"
	"context"
	"gorm.io/gorm"
)

//User数据库操作对象
type UserDao struct {
	*gorm.DB
}

//这个需要改成单例模式吗
func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{
		NewDBClient(ctx),
	}
}

// CreateUser CRUD

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) error {
	db := dao.Model(&model.User{}).Create(&user)
	return db.Error
}

// FindUserByUserName 根据用户名查找用户
func (dao *UserDao) FindUserByUserName(userName string) (*model.User, error) {
	res := &model.User{}
	db := dao.Model(&model.User{}).
		Where("user_name = ?", userName).Find(&res)
	return res, db.Error
}

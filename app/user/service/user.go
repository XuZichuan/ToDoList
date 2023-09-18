package service

import (
	"ToDoList/app/user/repo/db/dao"
	"ToDoList/app/user/repo/db/model"
	"ToDoList/consts"
	"ToDoList/idl/pb"
	"context"
	"errors"
	"sync"
)

var UserSvcIns *UserSvc
var UserSvcOnce sync.Once

// UserSvc 实现pb文件中定义的方法接口。此时就可以远程调用了。
type UserSvc struct {
}

// GetUserSvc 实例化用户svc对象，用于handler、rpc等调用
func GetUserSvc() *UserSvc {
	UserSvcOnce.Do(func() {
		UserSvcIns = &UserSvc{}
	})
	return UserSvcIns
}

// UserLogin 逻辑处理CRUD
func (s *UserSvc) UserLogin(ctx context.Context, req *pb.UserRequest, resp *pb.UserDetailResponse) error {
	resp.Code = consts.LoginSuccess
	//登陆逻辑 --查看此人是否存在，
	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		resp.Code = consts.LoginFail
		return errors.New("用户不存在")
	}
	// 数据库中拿到的密码，和用户输入的密码进行比对
	if user.CheckPassWord(req.Password) {
		resp.Code = consts.LoginFail
		return errors.New("密码错误")
	}
	// 登陆成功
	resp.UserDetail = GetPBUserModel(user)
	return nil
}
func (s *UserSvc) UserRegister(ctx context.Context, req *pb.UserRequest, resp *pb.UserDetailResponse) error {
	resp.Code = consts.LoginSuccess
	//注册时先校验两次密码是否一致
	if req.Password != req.PasswordConfirm {
		resp.Code = consts.LoginFail
		return errors.New("两次密码输入不一致")
	}
	//校验用户是否已存在
	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		resp.Code = consts.LoginFail
		return errors.New("获取用户失败")
	}
	//拿到了一个已存在的用户,通过主键来区别，防止因为数据库出错导致异常信息不对。
	if user.ID > 1 {
		resp.Code = consts.LoginFail
		return errors.New("该用户名已存在")
	}
	//校验成功,开始创建账号流程
	user = &model.User{
		UserName: req.UserName,
	}
	//密码加密
	err = user.SetPassWord(req.Password)
	if err != nil {
		resp.Code = consts.LoginFail
		return errors.New("密码加密失败")
	}
	//创建用户 把用户存到库里
	err = dao.NewUserDao(ctx).CreateUser(user)
	if err != nil {
		resp.Code = consts.LoginFail
		return errors.New("创建用户失败")
	}
	//注册成功
	return nil
}

// BuildUser 转换为pb文件中定义的UserModel
func GetPBUserModel(user *model.User) *pb.UserModel {
	return &pb.UserModel{
		Id:        uint32(user.ID),
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}
}

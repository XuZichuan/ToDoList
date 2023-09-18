package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const PassWordCost = 9

//user 数据库中的元对象
type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	PassWord string
}

// SetPassWord 注册用户后需要先加密密码，再设置给用户，后续存入数据库中
func (u *User) SetPassWord(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.PassWord = string(bytes)
	return nil
}

// CheckPassWord 检查用户密码
func (u *User) CheckPassWord(password string) bool {
	//利用bcrypt比较密码是否相同 u.PassWord应该是加密后的、password应该是前段传过来没加密的
	err := bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(password))
	//如果有err，则说明密码不相同
	return err == nil
}

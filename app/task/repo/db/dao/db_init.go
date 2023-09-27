package dao

import (
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"ToDoList/app/task/repo/db/model"
	"ToDoList/config"
)

// 用于初始化连接DB

var dbInstance *gorm.DB

// InitDB 连接数据库DB
func InitDB() {
	var gormLogger logger.Interface
	if gin.Mode() == "debug" {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default
	}
	//利用mysql的driver打开并连接数据库。
	curDB, err := gorm.Open(mysql.New(mysql.Config{ //mysql.New 返还一个连接器
		DSN:                       getDSN(), //获取dsn连接字
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据版本自动配置
	}), &gorm.Config{ //设置gorm的配置
		Logger: gormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	dbInstance = curDB
	//自动创建表
	dbInstance.Set(`gorm:table_options`, "charset=utf8mb4").AutoMigrate(&model.Task{})
}

func getDSN() string {
	//将配置文件中的数据读过来，组装成MySQL的连接字
	host := config.DbHost
	port := config.DbPort
	database := config.DbName
	username := config.DbUser
	password := config.DbPassWord
	charset := config.Charset
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")", "/", database, "?charset=" + charset + "&parseTime=true"}, "")
	return dsn
}

func NewDBClient(ctx context.Context) *gorm.DB {
	return dbInstance.WithContext(ctx)
}

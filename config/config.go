package config

import (
	"gopkg.in/ini.v1"

	"log"
)

var (
	//数据库配置
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	Charset    string
	//ETCD配置
	EtcdHost string
	EtcdPort string

	UserServiceAddress string
	TaskServiceAddress string
	HttpServiceAddress string
)

//用于初始化读取配置文件、包含MYSQL 中间件等
func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Println("配置文件加载失败")
	}
	LoadMySQLConfig(file)
	LoadETCDConfig(file)
	LoadServerConfig(file)
}

//从配置文件中把数据读到内存里。
func LoadMySQLConfig(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
	Charset = file.Section("mysql").Key("Charset").String()
}

func LoadETCDConfig(file *ini.File) {
	EtcdHost = file.Section("etcd").Key("EtcdHost").String()
	EtcdPort = file.Section("etcd").Key("EtcdPort").String()
}
func LoadServerConfig(file *ini.File) {
	UserServiceAddress = file.Section("server").Key("UserServiceAddress").String()
	TaskServiceAddress = file.Section("server").Key("TaskServiceAddress").String()
	HttpServiceAddress = file.Section("server").Key("HttpServiceAddress").String()
}

package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"mall/dao"
	"strings"
)

var (
	AppMode     string
	HttpPort    string
	UploadModel string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func Init() {

	// 从本地读取环境变量

	file, err := ini.Load("./conf/conf.ini")

	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}

	loadServer(file)
	loadMysqlData(file)
	//loadRedisData(file)

	// mysql 读
	connReadPath := strings.Join([]string{
		DbUser,
		":",
		DbPassWord,
		"@tcp(", DbHost, ":", DbPort, ")/",
		DbName,
		"?charset=utf8&parseTime=true"},
		"",
	)
	// mysql 写
	connWritePath := strings.Join([]string{
		DbUser,
		":",
		DbPassWord,
		"@tcp(", DbHost, ":", DbPort, ")/",
		DbName,
		"?charset=utf8&parseTime=true"},
		"",
	)

	dao.Database(connReadPath, connWritePath)
}

func loadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func loadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func loadRedisData(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

package config

import "github.com/spf13/viper"

/**
配置项
*/
type Config struct {
	DB *DbConfig
	ES *ESConfig
}

/**
数据库配置
*/
type DbConfig struct {
	Dns      string // 	基础连接结构
	Host     string // 	数据库host
	Port     string // 	数据库端口
	Database string // 	用于链接存储的库
	User     string //	用户名
	PassWord string //	密码
	Plugin   string //	数据解析插件
	Slot     string
}

/**
通用配置项
*/
type ESConfig struct {
	HttpPort   string // api服务端口
	RpcPort    string // rpc服务端口
	Namespace  string // 命名空间名称，用于确认该服务所在的datrix所属的租户
	EaglesIp   string // es的连接地址
	EaglesPort string // es连接地址的端口
}

/**
获取配置信息
*/
func GetConfig() *Config {
	// 设置从环境变量读取
	viper.AutomaticEnv()
	return &Config{
		DB: getDbConfig(),
		ES: getESConfig(),
	}
}

package config

import (
	"github.com/spf13/viper"
)

/**
获取常规配置信息
如端口，日志等级等
*/
func getESConfig() *ESConfig {
	return &ESConfig{
		HttpPort:   "9875", //暂时写死，不读环境变量
		RpcPort:    "9876",
		Namespace:  viper.GetString("Namespace"),
		EaglesIp:   "192.168.51.203",
		EaglesPort: "17100",
	}
}

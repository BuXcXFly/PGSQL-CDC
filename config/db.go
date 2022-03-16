package config

/**
获取数据库配置信息
*/
func getDbConfig() *DbConfig {
	return &DbConfig{
		Dns:      "postgres://%s:%s@%s:%s/%s?application_name=cdc", // postgres://库名:库密码@库连接地址/表名?application_name=cdc
		Host:     "192.168.51.203",                                 // todo 改成环境变量 viper.GetString("DbHost"),
		Port:     "30238",                                          // todo 改成环境变量 viper.GetString("DbPort"),
		Database: "asset",                                          // todo 改成环境变量 viper.GetString("DataBase"),
		User:     "datatom",                                        // todo 改成环境变量 viper.GetString("DbUser"),
		PassWord: "db_password",                                    // todo 改成环境变量 viper.GetString("DbPwd"),
		Plugin:   "test_decoding",
		Slot:     "slot_test",
	}
}

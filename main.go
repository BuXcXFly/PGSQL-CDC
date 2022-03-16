package main

import (
	"fmt"
	"resource-CDC/config"
	"resource-CDC/logic"
	"time"
)

func main() {
	// postgres://库名:库密码@库连接地址/表名?application_name=cdc
	// dsn := "postgres://datatom:db_password@192.168.51.203:30238/asset?application_name=cdc"
	// plugin := "test_decoding"
	// slot := "slot_test"
	// // if len(os.Args) > 1 {
	// // 	dsn = os.Args[1]
	// // }
	// // if len(os.Args) > 2 {
	// // 	plugin = os.Args[2]
	// // }
	// // if len(os.Args) > 3 {
	// // 	slot = os.Args[3]
	// // }
	// subscriber := &logic.Subscriber{
	// 	URL:    dsn,
	// 	Slot:   slot,
	// 	Plugin: plugin,
	// }

	conf := config.GetConfig()
	subscriber := &logic.Subscriber{
		URL:    fmt.Sprintf(conf.DB.Dns, conf.DB.User, conf.DB.PassWord, conf.DB.Host, conf.DB.Port, conf.DB.Database),
		Slot:   conf.DB.Slot,
		Plugin: conf.DB.Plugin,
	} // 创建新的CDC客户端
	// subscriber.DropReplicationSlot() // 如果存在，清理掉遗留的Slot

	subscriber.Connect() // 建立复制连接
	// defer subscriber.DropReplicationSlot() // 程序中止前清理掉复制槽
	subscriber.CreateReplicationSlot() // 创建复制槽(若已创建，会自动连接原有的复制槽中)
	subscriber.StartReplication()      // 开始接收变更流
	go func() {
		for {
			time.Sleep(5 * time.Second)
			subscriber.ReportProgress()
		}
	}() // 协程2每5秒地向主库汇报进度
	subscriber.Subscribe() // 主消息循环

	//初始化Es的客户端
	// if err := connect.InitEsClient(conf.ES.Namespace, conf.ES.HttpPort); err != nil {
	// 	fmt.Printf("InitEsClient：%s", err.Error())
	// 	os.Exit(1)
	// }

}

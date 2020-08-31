package global

import (
	"log"

	"github.com/yunshu2009/blog-service/pkg/setting"
)

// 定义包全局变量
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *log.Logger
)

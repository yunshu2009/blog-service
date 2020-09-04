package global

import (
	"github.com/yunshu2009/blog-service/pkg/setting"
	"github.com/yunshu2009/blog-service/pkg/logger"
)

// 定义包全局变量
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)

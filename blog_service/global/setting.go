package global

import (
	"GoTour/blog_service/pkg/logger"
	"GoTour/blog_service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings

	Logger *logger.Logger
)

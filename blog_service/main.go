// @title 博客系统
// @version 1.0
// @description 自学《一起用go做项目》

package main

import (
	"GoTour/blog_service/global"
	"GoTour/blog_service/internal/model"
	"GoTour/blog_service/internal/routers"
	"GoTour/blog_service/pkg/logger"
	"GoTour/blog_service/pkg/setting"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupSettings()
	if err != nil {
		log.Fatalf("main.init.setupSettings err:%v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("main.init.setupDBEngine err:%v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("main.init.setupLogger err:%v", err)
	}
}

func main() {
	//global.Logger.Infof("%s %s", "zfq", "go tour")
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeOut,
		WriteTimeout:   global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSettings() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEgine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	filename := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  filename,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

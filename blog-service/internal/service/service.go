package service

import (
	"github.com/aqin97/GoTour/blog-service/global"
	"github.com/aqin97/GoTour/blog-service/internal/dao"
	"github.com/gin-gonic/gin"
)

type Service struct {
	ctx *gin.Context
	dao *dao.Dao
}

func New(ctx *gin.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}

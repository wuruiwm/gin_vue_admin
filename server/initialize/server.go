package initialize

import (
	"app/api/middleware"
	"app/api/response"
	"app/global"
	"app/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

//Server 初始化gin的http服务
func Server() {
	//是否开启debug
	if global.Config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	//实例化gin
	r := gin.New()
	//Logger中间件
	r.Use(middleware.Logger)
	//Recovery中间件
	r.Use(middleware.Recover)
	//上传文件目录
	r.Static("/upload", "./resource/upload")
	//处理不存在的路由
	r.NoMethod(ErrRoute)
	r.NoRoute(ErrRoute)
	//路由设置
	r = router.Router(r)
	//启动服务
	_ = r.Run(fmt.Sprintf("0.0.0.0:%d", global.Config.ServerPort))
}

func ErrRoute(c *gin.Context) {
	response.Error(c, fmt.Sprintf("api不存在 path:%s method:%s", c.Request.URL, c.Request.Method))
}

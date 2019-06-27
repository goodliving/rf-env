package routers

import (

	"github.com/gin-gonic/gin"
	"goodliving/rf-env/routers/api/v1"
	"goodliving/rf-env/pkg/setting"
)

func InitRouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    gin.SetMode(setting.RunMode)
    apiv1 := r.Group("/api/v1")
    {
        //获取环境信息
        apiv1.GET("/env/:env", v1.GetEnv)
    }
    return r
}
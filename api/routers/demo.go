package routers

import (
	"github.com/Madou-Shinni/lottery-draw/api/handle"
	"github.com/Madou-Shinni/lottery-draw/internal/data"
	"github.com/Madou-Shinni/lottery-draw/internal/service"
	"github.com/gin-gonic/gin"
)

// 注册路由
func DemoRouterRegister(r *gin.Engine) {
	demoGroup := r.Group("demo")
	demoHandle := handle.NewDemoHandle(service.NewDemoService(data.NewDemoRepo()))
	{
		demoGroup.POST("", demoHandle.Add)
		demoGroup.DELETE("", demoHandle.Delete)
		demoGroup.DELETE("/delete-batch", demoHandle.DeleteByIds)
		demoGroup.GET("", demoHandle.Find)
		demoGroup.GET("/list", demoHandle.List)
		demoGroup.PUT("", demoHandle.Update)
	}
}

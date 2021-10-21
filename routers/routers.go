package routers

import (
	"github.com/gin-gonic/gin"
	"test/controller"
	"test/logger"
)

func Setuprouter() (r *gin.Engine) {
	r = gin.New()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))

/*	//这个是要有前端的模板  我用postman来测试
	r.Static("/static","./static")
	r.LoadHTMLGlob("./template/*")

	//如果确定不用往里面传参数就不用打括号
	r.GET("/index", controller.IndexHander)
*/

	grouptest := r.Group("/v1")
	{
		//提交数据
		grouptest.POST("/todo",controller.GetDate )

		//展示表里面的信息
		grouptest.GET("/todo", controller.ShowallDate)

		//更新数据
		//你在打网址传值的时候不用打： 直接打id就好了
		grouptest.PUT("/todo/:id", controller.UpdateaDate)

		//删除数据
		grouptest.DELETE("/todo/:id",controller.DeleteaDate)
	}
	return
}

package routers

import (
	"example/middleware/jwt"
	v1 "example/routers/api/v1"
	"github.com/gin-gonic/gin"
)
func InitRouter()*gin.Engine  {
	r:=gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode("debug")
	//日志中间件
	//r.Use(logger.LoggerToFile())

	r.POST("/auth",v1.Auth)


	apiv1:=r.Group("api/v1")

	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags",v1.GetTags)
		apiv1.POST("/tags",v1.AddTags)
		apiv1.PUT("/tags",v1.EditTag)
		apiv1.DELETE("/tags",v1.DeleteTag)

		apiv1.GET("/articles",v1.GetArticle)
		apiv1.POST("/articles",v1.AddArticle)
		apiv1.PUT("/articles",v1.EditArticle)
		apiv1.DELETE("/articles",v1.DeleteArticle)


	}

	r.GET("/export",v1.ExportTag)
	r.GET("/email",v1.SendEmail)
	r.GET("/qrcode",v1.Qrcode)
	r.GET("/poster",v1.Poster)


	return r
}
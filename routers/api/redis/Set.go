package redis

import (
	"github.com/gin-gonic/gin"
	"example/redis"
	"net/http"
)

//操作redis(写入)
func Set(c *gin.Context){
	a := c.DefaultQuery("c","NIUBI")
	//操作Redis
	if err:=redis.Set(a); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":http.StatusBadRequest,
			"msg":err.Error(),
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"msg":"成功",
	})
}
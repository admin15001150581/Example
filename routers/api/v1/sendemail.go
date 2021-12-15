package v1

import (
	"github.com/gin-gonic/gin"
	"example/pkg/e"
	"example/pkg/email"
	"net/http"
)

func SendEmail(c *gin.Context){
	email.SendEmail()
	c.JSON(http.StatusOK,gin.H{
		"code":e.SUCCESS,
		"msg":e.GetMsg(e.SUCCESS),
	})
}
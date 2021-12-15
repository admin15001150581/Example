package v1

import (
	"github.com/gin-gonic/gin"
	"example/models"
	"example/pkg/e"
	"example/pkg/until"
	"example/pkg/validator"
	"net/http"
)

var token string
var err error
func Auth(c *gin.Context){

	var code int
	var msg string
	var auth validator.Auth
	if err=c.ShouldBind(&auth); err!=nil{//参数校验
		code = e.INVALID_PARAMS
		msg = e.GetMsg(e.INVALID_PARAMS)
	}

	if models.CheckAuth(auth.Username,auth.Password){
		token ,err= until.GenerateToken(auth.Username,auth.Password)
		if err!=nil{
			code = e.INVALID_PARAMS
			msg = e.GetMsg(e.INVALID_PARAMS)
		}else{
			code = e.SUCCESS
			msg = e.GetMsg(e.SUCCESS)
		}

	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
		"data":token,
	})

}
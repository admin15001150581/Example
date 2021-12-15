package jwt

import (
	"github.com/gin-gonic/gin"
	"example/pkg/e"
	"example/pkg/until"
	"net/http"
	"time"
)

//建立中间件
func JWT() gin.HandlerFunc {
	return func (c *gin.Context){
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")

		if token ==""{
			code = e.INVALID_PARAMS
		}else{
			claims,err:=until.ParseToken(token)
			if err !=nil{
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}else if time.Now().Unix()>claims.ExpiresAt{ //token过期
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
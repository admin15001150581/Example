package v1

import (
	"github.com/gin-gonic/gin"
	"image/png"
	"net/http"
	q "example/pkg/qrcode"
	"example/pkg/e"
	"os"
)



func Qrcode(c *gin.Context){

	var (
		code int
		msg string

	)
	if img ,err :=q.CreateAvatar(); err!=nil{
		code=e.INVALID_PARAMS
		msg=err.Error()
	}else {
		code = e.SUCCESS
		msg = e.GetMsg(e.SUCCESS)
		f, _ := os.Create("image.png")
		png.Encode(f, img)
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
	})
}
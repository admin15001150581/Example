package v1

import (
	"github.com/gin-gonic/gin"
	"example/pkg/e"
	q "example/pkg/qrcode"
	"image/png"
	"net/http"
	"os"
)

func Poster(c *gin.Context)  {
	var (
		code int
		msg string

	)
	if  img,err:=q.Poster(); err!=nil{
		code=e.INVALID_PARAMS
		msg=err.Error()
	}else {
		code = e.SUCCESS
		msg = e.GetMsg(e.SUCCESS)
		f, _ := os.Create("image1.png")
		png.Encode(f, img)
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
	})
}
package v1

import (
	"example/models"
	"example/pkg/e"
	"example/pkg/excel"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	code int
	msg string
)


func ExportTag(c *gin.Context)  {
	datas := make(map[string]interface{})
	zhi:=models.GetTag(datas)
	if err:=excel.ExportExcel(zhi);  err!=nil{//导出失败
			code=e.ERROR
			msg=err.Error()

	}else{
		//下载
		c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "Book1.xlsx")) //指定下载时的默认文件名
		c.Writer.Header().Set("Content-Type", "octet-stream")
		c.File("./Book1.xlsx")

		code = e.SUCCESS
		msg = e.GetMsg(e.SUCCESS)
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
	})
}
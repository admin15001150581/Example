package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"example/models"
	"example/pkg/e"
	"example/pkg/validator"
	"net/http"
)

func GetTags(c *gin.Context)  {
	name:=c.Query("name")

	maps := make(map[string]interface{})
	datas:= make(map[string]interface{})

	if name!=""{
		maps["name"] = name
	}

	state:=-1

	if arg:=c.Query("state");arg!=""{
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	datas["list"] = models.GetTag(maps)

	c.JSON(http.StatusOK,gin.H{
		"code":e.SUCCESS,
		"msg":datas,
	})

}

func AddTags(c *gin.Context){

	var code int
	var msg string
	var tag validator.AddTag

	if err:=c.ShouldBind(&tag);err!=nil{//参数有误
		code =e.INVALID_PARAMS
		msg =err.Error()
	}else {
		if !models.ExitTagByName(tag.Name){ //tag已存在
		   code = e.ERROR_EXIST_TAG
		   msg = e.GetMsg(e.ERROR_EXIST_TAG)
		}else {
			 models.AddTags(tag.Name,tag.CreatedBy,tag.State)
			 code = e.SUCCESS
			 msg = e.GetMsg(e.SUCCESS)

		}

	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
	})
}

func EditTag(c *gin.Context)  {
	var code int
	var msg string
	var tag validator.EditTag
	if err:=c.ShouldBind(&tag);err!=nil{
		code =e.INVALID_PARAMS
		msg =err.Error()
	}else{
		if models.ExitTagById(tag.ID){ //id存在
			models.EditTag(tag.ID,tag.Name,tag.ModifiedBy,tag.State)
			code =e.SUCCESS
			msg = e.GetMsg(e.SUCCESS)
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
	})

}


func DeleteTag(c *gin.Context){
	var code int
	var msg string
	if arg:=c.PostForm("id"); arg==""{
		code = e.INVALID_PARAMS
		msg = e.GetMsg(e.INVALID_PARAMS)
	}else{
		id:=com.StrTo(arg).MustInt()
		if models.ExitTagById(id){
			models.DeleteTag(id)
			code = e.SUCCESS
			msg = e.GetMsg(e.SUCCESS)
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
	})
}
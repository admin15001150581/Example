package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"example/models"
	"example/pkg/e"
	"example/pkg/validator"
	"net/http"
)

func GetArticle(c *gin.Context)  {
	maps:=make(map[string]interface{})
	datas:=make(map[string]interface{})

	name:=c.Query("name")
	if name!=""{
		maps["name"] = name
	}

	datas["list"] = models.GetArticle(maps)
	c.JSON(http.StatusOK,gin.H{
		"status":e.SUCCESS,
		"msg":datas,
	})
}

func AddArticle(c *gin.Context){

	var code int
	var msg string
	var article validator.AddArticle

	if err:=c.ShouldBind(&article);err!=nil{
		code = e.INVALID_PARAMS
		msg = err.Error()
	}else {
		if !models.ExitArticleByName(article.Title){
			models.AddArticle(article.Title,article.Desc,article.Content,article.CreatedBy,article.State,article.TagId)
			code = e.SUCCESS
			msg = e.GetMsg(e.SUCCESS)
		}else{
			code = e.ERROR_EXIST_ARTICLE
			msg = e.GetMsg(e.ERROR_EXIST_ARTICLE)
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
	})
}

func EditArticle(c *gin.Context){
	var code int
	var msg string
	var article validator.EditArticle

	if err:=c.ShouldBind(&article);err!=nil {
		code = e.INVALID_PARAMS
		msg = err.Error()
	}else{
		if models.ExitArticleById(article.ID){
			models.EditArticle(article.ID,article.Title,article.Desc,article.Content,article.ModifiedBy,article.State,article.TagId)
			code = e.SUCCESS
			msg = e.GetMsg(e.SUCCESS)
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
	})
}


func  DeleteArticle(c *gin.Context){
	var code int
	var msg string

	id:=com.StrTo(c.PostForm("id")).MustInt()
	if id!=0{
		if models.ExitArticleById(id){
			models.DelteArticle(id)
			code = e.SUCCESS
			msg = e.GetMsg(e.SUCCESS)
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
	})
}
package models

type Article struct {
	Model
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
	TagId int `json:"tag_id"`
	Tag Tag
}

func GetArticle(maps interface{})(article []Article){
	db.Preload("Tag").Where(maps).Find(&article)
	return
}

func ExitArticleByName(title string) bool{
	var  article Article
	db.Where("title=?",title).First(&article)
	if article.ID>0{
		return true
	}
	return false
}

func AddArticle(title, desc, content, created_by string, state,tag_id int)bool{
	db.Create(&Article{
		Title: title,
		Desc: desc,
		Content: content,
		CreatedBy: created_by,
		State: state,
		TagId: tag_id,
	})
	return true
}


func ExitArticleById(id int)bool{
	var article Article
	db.Where("id=?",id).First(&article)
	if article.ID>0{
		return true
	}
	return false
}

func EditArticle(id int,title,desc,content, modified_by string,state,tag_id int) bool{
	db.Model(&Article{}).Where("id=?",id).Update(Article{
		Title: title,
		Desc: desc,
		Content: content,
		ModifiedBy: modified_by,
		State: state,
		TagId: tag_id,
	})
	return  true
}


func DelteArticle(id int)bool{
	db.Where("id=?",id).Delete(&Article{})
	return true
}
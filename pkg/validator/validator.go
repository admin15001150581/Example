package validator

type Tag struct {
		ID int `form:"id" binding:"required"`
		Name  string`form:"name" binding:"required"`
		CreatedBy string `form:"created_by" binding:"required"`
		ModifiedBy string `form:"modified_by" binding:"required"`
		State int `form:"state" binding:"required"`
}

type AddTag struct {
	Name  string`form:"name" binding:"required"`
	CreatedBy string `form:"created_by" binding:"required"`
	State int `form:"state" binding:"required"`
}

type EditTag struct {
	ID int `form:"id" binding:"required"`
	Name  string`form:"name" binding:"required"`
	ModifiedBy string `form:"modified_by" binding:"required"`
	State int `form:"state" binding:"required"`
}

type AddArticle struct {
	Title string `form:"title"  binding:"required"`
	Desc string `form:"desc"  binding:"required"`
	Content string `form:"content"  binding:"required"`
	CreatedBy string `form:"created_by" binding:"required"`
	State int `form:"state" binding:"required"`
	TagId int `form:"tag_id" binding:"required"`
}

type EditArticle struct {
	ID int `form:"id" binding:"required"`
	Title string `form:"title"  binding:"required"`
	Desc string `form:"desc"  binding:"required"`
	Content string `form:"content"  binding:"required"`
	ModifiedBy string `form:"modified_by" binding:"required"`
	State int `form:"state" binding:"required"`
	TagId int `form:"tag_id" binding:"required"`
}

type Auth struct {
	Username string `form:"username" binding:"required"'`
	Password string `form:"password" binding:"required"`
}
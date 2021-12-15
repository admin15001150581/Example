package models

type Tag struct {
	Model
	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTag(maps interface{})(tag []Tag){
	db.Where(maps).Find(&tag)
	return
}

func ExitTagByName(name string)bool{
	var tag Tag
	db.Where("name=?",name).First(&tag)
	if tag.ID>0{
		return false
	}
	return true
}

func AddTags(name, created_by string,state int)bool{

	db.Create(&Tag{
		Name: name,
		CreatedBy: created_by,
		State: state,
	})
	return true
}

func ExitTagById(id int)bool{
	var tag Tag
	db.Where("id=?",id).First(&tag)

	if tag.ID>0{
		return true
	}
	return false
}

func EditTag(id int,name, modified_by string,state int) bool{
	tag:=Tag{
		Name: name,
		ModifiedBy: modified_by,
		State: state,
	}
	db.Model(Tag{}).Where("id=?",id).Update(&tag)
	return true
}

func DeleteTag(id int) bool{

	db.Delete(&Tag{},id)
	return true
}


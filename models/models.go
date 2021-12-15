package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"example/pkg/setting"
	"log"
)

type Model struct {
	ID int `json:"id"`
	CreatedOn int `json:"created_by"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn int `json:"deleted_on"`
}

 var db *gorm.DB

func init(){
	var (
		err error
		dbType, dbName, host, user, password, tablePrefix string
	)

	dbName = setting.Viper.GetString("mysql.database")
	dbType ="mysql"
	host = setting.Viper.GetString("mysql.host")
	user = setting.Viper.GetString("mysql.user")
	password = setting.Viper.GetString("mysql.password")
	tablePrefix = setting.Viper.GetString("mysql.tablePrefix")


	db,err = gorm.Open(dbType,fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",user,password,host,dbName))
	if err!=nil{
		log.Fatal(err)
		Close()
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix+defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

}

func Close()  {
	defer db.Close()
}
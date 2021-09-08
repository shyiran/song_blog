package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

func Init() {
	dbhost, _ := web.AppConfig.String("dbhost")
	dbport, _ := web.AppConfig.String("dbport")
	dbuser, _ := web.AppConfig.String("dbuser")
	dbpassword, _ := web.AppConfig.String("dbpassword")
	dbname, _ := web.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(User), new(Category), new(Post), new(Config), new(Comment))
}

//返回带前缀的表名
func TableName(str string) string {
	dbprefix, _ := web.AppConfig.String("dbprefix")
	return dbprefix + str
}

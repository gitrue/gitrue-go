package models

import "github.com/astaxie/beego/orm"

type Card struct {
	Id   int
	Title string `orm:"size(100)"`
	Context string `orm:"type(text)"`
	Tag string
	User *User `orm:"rel(fk)"`    //设置一对多关系
}

func init() {
	orm.RegisterModel(new(Card))
}

package models

import "github.com/astaxie/beego/orm"

type Card struct {
	Id   int
	Title string `orm:"size(100)" json:"title"`
	LikeCount int `orm:"size(6);default(1)" json:"likeCount"`
	ClickCount int `orm:"size(6);default(1)" json:"clickCount"`
	Context string `orm:"type(text)" json:"context"`
	Tag string `orm:"size(100);null" json:"tag"`
	User *User `orm:"rel(fk)"`    //设置一对多关系
	Comment []*Comment `orm:"reverse(many);null" json:"comments"` // 设置一对多的反向关系
}

func AddCard(card *Card) Card {
	o := orm.NewOrm()
	o.Using("default")
	o.Insert(card)//插入数据库

	return *card
}


func init() {
	orm.RegisterModel(new(Card))
}

package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Comment struct {
	Id   int
	Context string `orm:"size(1000)" json:"context"`
	LikeCount int `orm:"size(6);default(1)" json:"likeCount"`
	ClickCount int `orm:"size(6);default(1)" json:"clickCount"`
	NickName string `orm:"size(10)" json:"nikeName"`
	No int  `orm:"size(6)" json:"no"`
	Created time.Time `orm:"auto_now_add;type(datetime)" json:"create"`
	Card *Card `orm:"rel(fk)"`    //设置一对多关系
	Article *Article `orm:"rel(fk)" json:"article"`    //设置一对多关系
	User *User `orm:"rel(fk)" json:"user"`    //设置一对多关系
}

func AddComment(comment *Comment) Comment {
	o := orm.NewOrm()
	o.Using("default")
	o.Insert(comment)//插入数据库

	return *comment
}

func init() {
	orm.RegisterModel(new(Comment))
}

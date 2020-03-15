package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

//文章表
type Article struct {
	Id   int64 `json:"id"`
	Tag string `orm:"size(100);null" json:"tag"`
	Content string `orm:"type(text)" json:"content"`
	Des string `orm:"size(200)" json:"des"` // 0 -> girl  1 -> boy
	Order int `orm:"size(5)"  json:"order"` // 0 -> girl  1 -> boy
	//Tag string
	Created time.Time `orm:"auto_now_add;type(datetime)" json:"create"`
	User *User `orm:"rel(fk)"`    //设置一对多关系
	Comment []*Comment `orm:"reverse(many);null" json:"comments"` // 设置一对多的反向关系
}

func GetArticle(id int) Article {
	o := orm.NewOrm()
	o.Using("default")
	var article Article
	_,err := o.QueryTable(new(Article)).Filter("id", id).All(&article)
	fmt.Printf("ERR: %v\n", err)
	return article
}

func AddArticle(article *Article) Article {
	user := new(User)
	user.Id = 21323
	article.User = user

	o := orm.NewOrm()
	o.Using("default")
	o.Insert(article)//插入数据库

	return *article
}

func init() {
	orm.RegisterModel(new(Article))
}

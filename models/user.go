package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户基本 model
type User struct {
	Id   int
	NickName string `orm:"size(30)"`
	UserName string `orm:"size(30)"`
	Sex int // 0 -> girl  1 -> boy
	Tag string
	Empiric int //经验
	Level int `orm:"default(1)"`
	Avatar string
	WxId string `orm:"null"`
	Phone string `orm:"size(11);null"`
	Email string `orm:"size(50);null"`
	Password string `orm:"null"`
	Online bool
	Private bool
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	LastLogin time.Time `orm:"auto_now_add;type(datetime)"`
	UserAction *UserAction `orm:"reverse(one)"` // 反向关联 mysql不建立字段 model 保留 struct
	Card []*Card `orm:"reverse(many);null" json:"cards"` // 设置一对多的反向关系
	Article []*Article `orm:"reverse(many);null" json:"articles"` // 设置一对多的反向关系
	Comment []*Comment `orm:"reverse(many);null" json:"comments"` // 设置一对多的反向关系
}

//用户行为 model
type UserAction struct {
	Id   int
	Type string
	level int //用户行为等级 由高到低
	RelevanceId int // 关联的业务ID
	Msg string `orm:"size(100)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	User *User `orm:"rel(one)"` // OneToOne relation 自定建立 $table_id
}



func init() {
	orm.RegisterModel(new(UserAction),new(User))
}


//todo userInfo
//todo userHistory
//todo userCollect
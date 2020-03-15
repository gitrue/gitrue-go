package models

import "time"

//好战分享表
type Setting struct {
	Id   int64
	Type string `orm:"size(10)"`
	Value string `orm:"size(500)"`
	Des string `orm:"size(200)"` // 0 -> girl  1 -> boy
	//Tag string
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}


package models

import (
	"time"
)

//好战分享表
type Share struct {
	Id   int
	Link string `orm:"size(50)"`
	Image string `orm:"size(100)"`
	des string `orm:"size(200)"` // 0 -> girl  1 -> boy
	Tag string
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}


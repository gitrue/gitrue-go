package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"sort"
	"time"
)



//好战分享表
type Setting struct {
	Id   int64 `json:"id"`
	Type string `orm:"size(10)" json:"type"`
	Value string `orm:"size(500)" json:"value"`
	Des string `orm:"size(200)" json:"des"` // 0 -> girl  1 -> boy
	Order int `orm:"size(5)"  json:"order"` // 0 -> girl  1 -> boy
	//Tag string
	Created time.Time `orm:"auto_now_add;type(datetime)" json:"create"`
}

type Settings []*Setting

// Len()方法和Swap()方法不用变化
// 获取此 slice 的长度
func (s Settings) Len() int { return len(s) }

// 交换数据
func (s Settings) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// 嵌套结构体  将继承 Person 的所有属性和方法
// 所以相当于SortByName 也实现了 Len() 和 Swap() 方法
type SortByOrder struct{ Settings }

// 根据元素的姓名长度降序排序 （此处按照自己的业务逻辑写）
func (p SortByOrder) Less(i, j int) bool {
	return p.Settings[i].Order < p.Settings[j].Order
}



func GetSetting(typeName string) []*Setting {
	o := orm.NewOrm()
	var settings Settings
	_,err := o.QueryTable(new(Setting)).Filter("type", typeName).All(&settings)
	sort.Sort(SortByOrder{settings})
	fmt.Printf("ERR: %v\n", err)
	return settings
}

func init() {
	orm.RegisterModel(new(Setting))
}

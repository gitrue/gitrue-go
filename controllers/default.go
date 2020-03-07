package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type Result struct {
	Code int `json:"code"` // 指定 json 小写 key
	Desc string `json:"desc"`
	Detail interface{} `json:"detail"`
}
func (c *MainController) Get() {
	// test json
	data := &Result{200, "获取成功","Hello world "}
	c.Data["json"] = data
	c.ServeJSON()
}

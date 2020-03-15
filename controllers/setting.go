package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gitrue-go/models"
)
var o = orm.NewOrm()
type SettingController struct {
	beego.Controller
}

// curl localhost:9999/api/setting/get/homeTags
// @router /get/:type [get]
func (c *SettingController) Get() {
	typeName := c.Ctx.Input.Param(":type")
	o := orm.NewOrm()
	var settings []orm.ParamsList
	_,err := o.QueryTable(new(models.Setting)).Filter("type", typeName).ValuesList(&settings)
	fmt.Printf("ERR: %v\n", err)
	// test json
	//data := &Result{200, "获取成功","Hello world "}
	c.Data["json"] = settings
	c.ServeJSON()
}

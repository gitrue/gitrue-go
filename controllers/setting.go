package controllers

import (
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
	c.Data["json"] = models.GetSetting(typeName)
	c.ServeJSON()
}

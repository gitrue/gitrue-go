package controllers

import (
	"encoding/json"
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
	c.Data["json"] = models.GetSetting(typeName)
	c.ServeJSON()
}

// @router /insert [post]
func (c *SettingController) Post() {
	var setting models.Setting
	json.Unmarshal(c.Ctx.Input.RequestBody, &setting)
	uid := models.AddSetting(&setting)
	c.Data["json"] = uid
	fmt.Println(uid)
	c.ServeJSON()
}
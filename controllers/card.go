package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"gitrue-go/models"
)
type CardController struct {
	beego.Controller
}

// curl localhost:9999/api/setting/get/homeTags
// @router /get/:type [get]
func (c *CardController) GetCard() {
	//typeName := c.Ctx.Input.Param(":type")
	//c.Data["json"] = models.GetSetting(typeName)
	//c.ServeJSON()
}

// @router /insert [post]
func (c *CardController) PostCard() {
	var card models.Card
	json.Unmarshal(c.Ctx.Input.RequestBody, &card)
	uid := models.AddCard(&card)
	c.Data["json"] = uid
	fmt.Println(uid)
	c.ServeJSON()
}
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"gitrue-go/models"
)
type CommentController struct {
	beego.Controller
}

// curl localhost:9999/api/setting/get/homeTags
// @router /get/:type [get]
func (c *CommentController) GetComment() {
	//typeName := c.Ctx.Input.Param(":type")
	//c.Data["json"] = models.GetSetting(typeName)
	//c.ServeJSON()
}

// @router /insert [post]
func (c *CommentController) PostComment() {
	var comment models.Comment
	json.Unmarshal(c.Ctx.Input.RequestBody, &comment)
	uid := models.AddComment(&comment)
	c.Data["json"] = uid
	fmt.Println(uid)
	c.ServeJSON()
}
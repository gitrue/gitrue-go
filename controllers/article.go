package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"gitrue-go/models"
)
type ArticleController struct {
	beego.Controller
}

// curl localhost:9999/api/setting/get/homeTags
// @router /get/:type [get]
func (c *ArticleController) GetArticle() {
	//typeName := c.Ctx.Input.Param(":type")
	//c.Data["json"] = models.a(typeName)
	//c.ServeJSON()
}

// @router /insert [post]
func (c *ArticleController) PostArticle() {
	var article models.Article
	json.Unmarshal(c.Ctx.Input.RequestBody, &article)
	uid := models.AddArticle(&article)
	c.Data["json"] = uid
	fmt.Println(uid)
	c.ServeJSON()
}
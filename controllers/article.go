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
// @router /get/:id [get]
func (c *ArticleController) GetArticle() {
	//id ,_:= c.GetInt(":id")
	c.Data["json"] = models.GetArticle(1)
	c.ServeJSON()
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
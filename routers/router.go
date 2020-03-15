package routers

import (
	"github.com/astaxie/beego"
	"gitrue-go/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/setting",
			beego.NSInclude(
				&controllers.SettingController{},
			),
		),
		beego.NSNamespace("/comment",
			beego.NSInclude(
				&controllers.CommentController{},
			),
		),
		beego.NSNamespace("/card",
			beego.NSInclude(
				&controllers.CardController{},
			),
		),
		beego.NSNamespace("/article",
			beego.NSInclude(
				&controllers.ArticleController{},
			),
		),

	)
	beego.AddNamespace(ns)
	//beego.Include(&controllers.SettingController{})
}

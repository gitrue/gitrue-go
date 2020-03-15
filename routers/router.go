package routers

import (
	"gitrue-go/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/setting",
			beego.NSInclude(
				&controllers.SettingController{},
			),
		),

	)
	beego.AddNamespace(ns)
	//beego.Include(&controllers.SettingController{})
}

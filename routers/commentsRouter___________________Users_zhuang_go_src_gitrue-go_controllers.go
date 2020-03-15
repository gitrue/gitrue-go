package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["gitrue-go/controllers:SettingController"] = append(beego.GlobalControllerRouter["gitrue-go/controllers:SettingController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/get/:type`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

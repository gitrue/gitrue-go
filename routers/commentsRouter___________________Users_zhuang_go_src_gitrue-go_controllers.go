package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["gitrue-go/controllers:ArticleController"] = append(beego.GlobalControllerRouter["gitrue-go/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetArticle",
            Router: `/get/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitrue-go/controllers:ArticleController"] = append(beego.GlobalControllerRouter["gitrue-go/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "PostArticle",
            Router: `/insert`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitrue-go/controllers:CardController"] = append(beego.GlobalControllerRouter["gitrue-go/controllers:CardController"],
        beego.ControllerComments{
            Method: "GetCard",
            Router: `/get/:type`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitrue-go/controllers:CardController"] = append(beego.GlobalControllerRouter["gitrue-go/controllers:CardController"],
        beego.ControllerComments{
            Method: "PostCard",
            Router: `/insert`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitrue-go/controllers:CommentController"] = append(beego.GlobalControllerRouter["gitrue-go/controllers:CommentController"],
        beego.ControllerComments{
            Method: "GetComment",
            Router: `/get/:type`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitrue-go/controllers:CommentController"] = append(beego.GlobalControllerRouter["gitrue-go/controllers:CommentController"],
        beego.ControllerComments{
            Method: "PostComment",
            Router: `/insert`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitrue-go/controllers:SettingController"] = append(beego.GlobalControllerRouter["gitrue-go/controllers:SettingController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/get/:type`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitrue-go/controllers:SettingController"] = append(beego.GlobalControllerRouter["gitrue-go/controllers:SettingController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/insert`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

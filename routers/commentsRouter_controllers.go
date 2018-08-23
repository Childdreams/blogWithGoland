package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["myproject/controllers:AdminController"] = append(beego.GlobalControllerRouter["myproject/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Blank",
			Router: `/blank`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myproject/controllers:AdminController"] = append(beego.GlobalControllerRouter["myproject/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Charts",
			Router: `/charts`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myproject/controllers:AdminController"] = append(beego.GlobalControllerRouter["myproject/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Error404",
			Router: `/error`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myproject/controllers:AdminController"] = append(beego.GlobalControllerRouter["myproject/controllers:AdminController"],
		beego.ControllerComments{
			Method: "ForgotPassword",
			Router: `/forgot-password`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myproject/controllers:AdminController"] = append(beego.GlobalControllerRouter["myproject/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Tables",
			Router: `/tables`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}

package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["Bee-api-Project/controllers:BannerControllers"] = append(beego.GlobalControllerRouter["Bee-api-Project/controllers:BannerControllers"],
		beego.ControllerComments{
			Method: "GetBanner",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"] = append(beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"] = append(beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"] = append(beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"] = append(beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"] = append(beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"] = append(beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"],
		beego.ControllerComments{
			Method: "CreatePost",
			Router: `/createpost`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"] = append(beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAllPost",
			Router: `/getpost`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"] = append(beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetPostFromId",
			Router: `/getpostfromid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"] = append(beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetPostFromTag",
			Router: `/getpostfromtagId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"] = append(beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAllTag",
			Router: `/gettag`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"] = append(beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"] = append(beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"] = append(beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"],
		beego.ControllerComments{
			Method: "ModifyPwd",
			Router: `/modifypwd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"] = append(beego.GlobalControllerRouter["Bee-api-Project/controllers:UserController"],
		beego.ControllerComments{
			Method: "RegisterUser",
			Router: `/regiser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}

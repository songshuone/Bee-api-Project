// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Bee-api-Project/controllers"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
	//"fmt"
	//"Bee-api-Project/models"
	"fmt"
	"crypto/md5"
	"math"
	"io"
	"github.com/astaxie/beego/context"
	"strings"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/banner",
			beego.NSInclude(
				&controllers.BannerControllers{},
			),
		),
	)
	ns.Filter("before", func(context *context.Context) {
		requestUrl:=context.Request.RequestURI
		if strings.Contains(requestUrl , "/v1/user/login" ){
			return
		}
		if  strings.Contains(requestUrl,"banner"){
			return
		}
		if  strings.Contains(requestUrl,"v1"){
			return
		}
		context.Request.ParseForm()
		session := context.Request.Form.Get("session")
		se := context.Input.Session("api")
		if se == nil || Md5(fmt.Sprint(se.(int))) != session {
			 re:= Response{Message:"未登录",Status:403}
			context.Output.JSON(re, true, true)
		}
	})
	beego.AddNamespace(ns)
}
func GetMD5(lurl string) string {
	h := md5.New()
	salt1 := "salt4shorturlwp123" + fmt.Sprint(math.Phi*math.Pi)
	io.WriteString(h, lurl+salt1)
	urlmd5 := fmt.Sprintf("%x", h.Sum(nil))
	return urlmd5
}

//md5加密        return 加密 后的字符串
//values  待加密的字符串
func Md5(values string) string {
	return GetMD5(values)
}
type Response struct {
	Message string `json:"message"`
	Status  int `json:"status"`
}

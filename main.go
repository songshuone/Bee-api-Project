package main

import (
	_ "Bee-api-Project/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/session"
)

var GlobalSessions  *session.Manager
func init() {

	//globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid",
	// "enableSetCookie,omitempty": true,
	// "gclifetime":3600,
	// "maxLifetime": 3600,
	// "secure": false,
	// "sessionIDHashFunc": "sha1",
	// "sessionIDHashKey": "",
	// "cookieLifeTime": 3600,
	// "providerConfig": ""}`)
	//go globalSessions.GC()
	GlobalSessions,_=session.NewManager("memory",&session.ManagerConfig{
		CookieName:"api",
		EnableSetCookie:true,
		Gclifetime:3600,
		Maxlifetime:3600,
		Secure:false,
		ProviderConfig:"",
		CookieLifeTime:3600,

	})
	defer globalSessions.GC()

	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/go?charset=utf8")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

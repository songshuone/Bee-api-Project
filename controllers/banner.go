package controllers

import (
	"github.com/astaxie/beego"
	"Bee-api-Project/models"
	"github.com/astaxie/beego/cache"
	"time"
)

var urlcache cache.Cache

type BannerControllers struct {
	beego.Controller
}

func init() {

	urlcache, _ = cache.NewCache("memory", `{"interval":0}`)
}

// URLMapping ...
func (c *BannerControllers) URLMapping() {
	c.Mapping("GetBanner", c.GetBanner)
}

//var responseData models.ResponseData
//
//func init() {
//	responseData.Result = ""
//	responseData.Status = 200
//	responseData.Message = ""
//}
// get ...
// @Title get
// @Description get banner
// @Success 200 {int} models.Banner
// @Failure 403 body is empty
// @router / [get]
func (c *BannerControllers) GetBanner() {
	if urlcache.IsExist("banner") {
		responseData.Result = urlcache.Get("banner")
		responseData.Message = "获取cache成功"
		c.Data["json"] = responseData
		c.ServeJSON()
		return
	}
	if banner, erro := models.GetBanner(); erro == nil {
		responseData.Status = 200
		responseData.Result = banner
		responseData.Message="获取数据成功"
		if len(banner)>0{
			urlcache.Put("banner",banner, 10 * time.Second)
		}
		c.Data["json"] = responseData
	} else {
		responseData.Status = 403
		responseData.Message = erro.Error()
		c.Data["json"] = responseData.Response
	}

	c.ServeJSON()
}

package controllers

import (
	"github.com/astaxie/beego"
	"Bee-api-Project/models"
)

type BannerControllers struct {
	beego.Controller
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
	if banner, erro := models.GetBanner(); erro == nil {
		responseData.Status = 200
		responseData.Result = banner
		c.Data["json"] = responseData
	} else {
		responseData.Status = 403
		responseData.Message = erro.Error()
		c.Data["json"] = responseData.Response
	}

	c.ServeJSON()
}

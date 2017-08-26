package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//auto_now / auto_now_add
//Created time.Time `orm:"auto_now_add;type(datetime)"`
//Updated time.Time `orm:"auto_now;type(datetime)"`
//auto_now 每次 model 保存时都会对时间自动更新
//auto_now_add 第一次保存时才设置时间
//对于批量的 update 此设置是不生效的
type Banner struct {
	Id         int64 `json:"id"`
	ImageUrl   string `orm:"null;size(60)" json:"image_url"`
	AdverUrl   string `orm:"null;size(60)" json:"adver_url"`
	BannerDesc string `orm:"null;size(10)" json:"banner_desc"`
	CreateDate time.Time `orm:"auto_now_add;type(datetime)" json:"create_date"`
	EndDate    time.Time `orm:"type(datetime)" json:"end_date"`
}

func init() {
	orm.RegisterModel(new(Banner))

}
func GetBanner() ([]*Banner,error) {
	o := orm.NewOrm()
	var banner []*Banner
	if _,erro:=o.QueryTable(new(Banner)).Limit(6, 0).All(&banner);erro!=nil{
		return nil,erro
	}else{
		return banner,nil
	}

}


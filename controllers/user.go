package controllers

import (
	"Bee-api-Project/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
	"fmt"
	"time"
	"github.com/DeanThompson/jpush-api-go-client/push"
	"github.com/DeanThompson/jpush-api-go-client"
)

// UserController operations for User
type UserController struct {
	beego.Controller
}

var responseData models.ResponseData

func init() {
	responseData.Result = ""
	responseData.Status = 200
	responseData.Message = ""
	jpushClient.SetDebug(true)
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Login", c.Login)
	c.Mapping("RegisterUser", c.RegisterUser)
	c.Mapping("ModifyPwd", c.ModifyPwd)
	c.Mapping("Logout", c.Logout)
	c.Mapping("GetAllPost", c.GetAllPost)
	c.Mapping("GetPostFromId", c.GetPostFromId)
	c.Mapping("GetPostFromTag", c.GetPostFromTag)
	c.Mapping("GetAllTag", c.GetAllTag)
	c.Mapping("CreatePost", c.CreatePost)
}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	var v models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddUser(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUserById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get User
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.User
// @Failure 403
// @router / [get]
func (c *UserController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10 //每页限制多少个数据
	var offset int64 = 0 //从那个位置开始查询

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}
	l, err := models.GetAllUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		responseData.Result = l
		responseData.Message = "获取数据成功"
		c.Data["json"] = responseData
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.User{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateUserById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}

	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUser(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

//Login...
//@Title Login
// @Description login the User
// @Param	name		query 	string	true		"The name you want to login"
// @Param	password		query 	string	true		"The password you want to login"
// @Success 200 {string} login success!
// @Failure 403 user no exist
//@router /login [get]
func (c *UserController) Login() {
	usernameStr := c.GetString("name")
	passwordStr := c.GetString("password")
	isSucess, u := models.Login(usernameStr, passwordStr)
	if isSucess {
		v := c.GetSession("api")
		var result models.UserLogin
		//r:=	rand.New(rand.NewSource(time.Now().UnixNano()))
		if v == nil {
			c.SetSession("api", time.Now().Nanosecond()+666)
		} else {
			c.SetSession("api", time.Now().Nanosecond()+666)
		}
		v = c.GetSession("api")

		Notifity("登陆成功", "温馨提示！")
		//test_Push()
		result.Session = models.Md5(fmt.Sprint(v.(int)))
		responseData.Message = "登录成功"
		responseData.Status = 200
		result.User = u
		responseData.Result = &result
		c.Data["json"] = responseData
	} else {
		responseData.Message = "用户名或密码不对"
		responseData.Status = 403
		c.Data["json"] = responseData.Response
	}
	c.ServeJSON()
}

//RegisterUser...
//@Title RegisterUser
// @Description login the User
// @Param	name		query 	string	true		"The name you want to login"
// @Param	password		query 	string	true		"The password you want to login"
// @Success 200 {string} RegisterUser success!
// @Failure 403 register fail
//@router /regiser [post]
func (u *UserController) RegisterUser() {
	name := u.GetString("name")
	password := u.GetString("password")
	responseData.Status = 403
	if len(name) < 3 || len(name) > 10 {
		responseData.Message = "用户名长度在3到10个"
		u.Data["json"] = responseData.Response
		u.ServeJSON()
		return
	}
	if len(password) < 3 || len(password) > 10 {
		responseData.Message = "密码长度在3到10个"
		u.Data["json"] = responseData.Response
		u.ServeJSON()
		return
	}
	erro := models.RegisterUser(name, password)
	if erro == nil {
		responseData.Message = "注册成功"
		responseData.Status = 200
		u.Data["json"] = responseData.Response
	} else {
		responseData.Message = erro.Error()
		responseData.Status = 403
		u.Data["json"] = responseData.Response
	}
	u.ServeJSON()
}

//ModifyPwd...
//@Title ModifyPwd
// @Description ModifyPwd the User
// @Param	id		query 	string	true		"The id you want to modify"
// @Param	session		query 	string	true		"The id you want to modify"
// @Param	password		query 	string	true		"The password you want to reset pwd"
// @Success 200 {string} 修改 success!
// @Failure 403 修改 fail
//@router /modifypwd [post]
func (u *UserController) ModifyPwd() {
	if err := CheckIsLogin(u); err != nil {
		responseData.Message = err.Error()
		responseData.Status = 403
		u.Data["json"] = responseData.Response
		u.ServeJSON()
		return
	}
	id := u.GetString("id")
	password := u.GetString("password")
	if len(id) == 0 {
		responseData.Message = "用户id有误"
		responseData.Status = 403
		u.Data["json"] = responseData.Response
		u.ServeJSON()
		return
	}
	if len(password) < 6 || len(password) > 15 {
		responseData.Message = "密码为6至15位"
		responseData.Status = 401
		u.Data["json"] = responseData.Response
		u.ServeJSON()
		return
	}
	if erro := models.ModifyPwd(id, password); erro == nil {
		v := u.GetSession("api")
		if v != nil {
			responseData.Message = "修改成功session;" + models.Md5(fmt.Sprint(v.(int)))

		} else {
			responseData.Message = "修改成功session=null"
		}
		responseData.Status = 200
		u.Data["json"] = responseData.Response
	} else {
		responseData.Message = erro.Error()
		responseData.Status = 401
		u.Data["json"] = responseData.Response
	}
	u.ServeJSON()
}

const (
	appKey       = "f7cc5881f5a0d47e4e077f8c"
	masterSecret = "cc9b1ac39511483e35592d9f"
)

var jpushClient = jpush.NewJPushClient(appKey, masterSecret)

/**
  通知到客户端
 */
func Notifity(messageNotification string, titleNotification string) {
	platForm := push.NewPlatform()
	platForm.Add("android") //设置指定平台
	audience := push.NewAudience()
	audience.All() //发送给所有人
	androidNotification := push.NewAndroidNotification(messageNotification)
	androidNotification.Title = titleNotification
	notification := push.NewNotification("Notification Alert")
	notification.Android = androidNotification
	payload := push.NewPushObject()
	payload.Platform = platForm
	payload.Audience = audience
	payload.Notification = notification
	result, err := jpushClient.Push(payload)
	if err != nil {
		fmt.Print("Push failed:", err)
	} else {
		fmt.Println("Push result:", result)
	}
}

//ModifyPwd...
//@Title ModifyPwd
// @Description ModifyPwd the User
// @Param	session		query 	string	true		"The id you want to modify"
// @Success 200 {string}注销成功
// @Failure 403 你还没有登录
//@router /logout [post]
func (c *UserController) Logout() {
	if CheckIsLogin(c) != nil {
		responseData.Message = "你还没有登录！"
		responseData.Status = 403
	} else {
		c.SetSession("api", nil)
		responseData.Message = "注销成功"
	}
	c.Data["json"] = responseData.Response
	c.ServeJSON()
}

//检查是否登录
func CheckIsLogin(c *UserController) error {
	session := c.GetString("session")
	se := c.GetSession("api")
	if se != nil && models.Md5(fmt.Sprint(se.(int))) == session {
		return nil
	}
	return errors.New("你还没有登录哦")

}

//getAllPost...
//@Title getAllPost
// @Description getAllPost the Post
// @Param	limit		query 	int	true		"The id you want to modify"
// @Param	offset		query 	int	true		"The id you want to modify"
// @Success 200 {string}获取数据成功
// @Failure 403 获取数据失败
//@router /getpost [get]
func (c *UserController) GetAllPost() {

	limit, erro := strconv.Atoi(c.GetString("limit"))
	offset, offseterro := strconv.Atoi(c.GetString("offset"))
	if erro != nil && offseterro != nil {
		responseData.Message = fmt.Sprint(erro, offseterro)
		responseData.Status = 403
		c.Data["json"] = responseData.Response
	} else {
		data, erro ,count:= models.GetAllPost(limit, offset)
		if erro != nil {
			responseData.Message = erro.Error()
			responseData.Status = 403
			c.Data["json"] = responseData.Response
		} else {
			responseData.Result = models.ResponseTotal{count,data}
			responseData.Status = 200
			responseData.Message = "获取数据成功"
			c.Data["json"] = responseData
		}
	}
	c.ServeJSON()
}

//GetPostFromId...
//@Title GetPostFromId
// @Description GetPostFromId the Post
// @Param	postId		query 	int	true		"The id you want to modify"
// @Success 200 {string}获取数据成功
// @Failure 403 获取数据失败
//@router /getpostfromid [get]
func (c *UserController) GetPostFromId() {
	postId, erro := strconv.Atoi(c.GetString("postId"))
	if erro != nil {
		responseData.Message = erro.Error()
		responseData.Status = 403
		c.Data["json"] = responseData.Response
	} else {
		post := models.GetPostFromId(postId)
		responseData.Status = 200
		responseData.Message = "获取数据成功"
		responseData.Result = post
		c.Data["json"] = responseData
	}
	c.ServeJSON()
}

//GetPostFromTag...
//@Title GetPostFromTag
// @Description GetPostFromTag the Post
// @Param	tagId		query 	int	true		"The id you want to modify"
// @Success 200 {string}获取数据成功
// @Failure 403 获取数据失败
//@router /getpostfromtagId [get]
func (c *UserController) GetPostFromTag() {
	tagId, err := strconv.Atoi(c.GetString("tagId"))
	if err != nil {
		responseData.Message = err.Error()
		responseData.Status = 403
		c.Data["json"] = responseData.Response
	} else {
		posts, erro := models.GetPostFromTag(tagId)
		if erro != nil {
			responseData.Message = erro.Error()
			responseData.Status = 403
			c.Data["json"] = responseData.Response
		} else {
			responseData.Message = "获取数据成功"
			responseData.Status = 200
			responseData.Result = posts
			c.Data["json"] = responseData
		}
	}
	c.ServeJSON()
}

//GetAllTag...
//@Title GetAllTag
// @Description GetAllTag the Post
// @Success 200 {string}获取数据成功
// @Failure 403 获取数据失败
//@router /gettag [get]
func (c *UserController) GetAllTag() {
	tags, err := models.GetAllTag()
	if err != nil {
		responseData.Status = 403
		responseData.Message = err.Error()
		c.Data["json"] = responseData.Response
	} else {
		responseData.Status = 200
		responseData.Message = "获取数据成功"
		responseData.Result = tags
		c.Data["json"] = responseData
	}
	c.ServeJSON()
}
func test_Push() {
	// platform 对象
	platform := push.NewPlatform()
	// 用 Add() 方法添加具体平台参数，可选: "all", "ios", "android"
	platform.Add("android")
	// 或者用 All() 方法设置所有平台
	// platform.All()

	// audience 对象，表示消息受众
	audience := push.NewAudience()
	//audience.SetTag([]string{"广州", "深圳"})   // 设置 tag
	//audience.SetTagAnd([]string{"北京", "女"}) // 设置 tag_and
	// 和 platform 一样，可以调用 All() 方法设置所有受众
	audience.All()

	// notification 对象，表示 通知，传递 alert 属性初始化
	notification := push.NewNotification("Notification Alert")

	// android 平台专有的 notification，用 alert 属性初始化
	androidNotification := push.NewAndroidNotification("Android Notification Alert")
	androidNotification.Title = "title"
	androidNotification.AddExtra("key", "value")

	notification.Android = androidNotification

	// iOS 平台专有的 notification，用 alert 属性初始化
	//iosNotification := push.NewIosNotification("iOS Notification Alert")
	//iosNotification.Badge = 1
	// Validate 方法可以验证 iOS notification 是否合法
	// 一般情况下，开发者不需要直接调用此方法，这个方法会在构造 PushObject 时自动调用
	// iosNotification.Validate()

	//notification.Ios = iosNotification

	// Windows Phone 平台专有的 notification，用 alert 属性初始化
	//wpNotification := push.NewWinphoneNotification("Winphone Notification Alert")
	// 所有平台的专有 notification 都有 AddExtra 方法，用于添加 extra 信息
	//wpNotification.AddExtra("key", "value")
	//wpNotification.AddExtra("extra_key", "extra_value")

	//notification.Winphone = wpNotification

	// message 对象，表示 透传消息，用 content 属性初始化
	message := push.NewMessage("Message Content must not be empty")
	message.Title = "Message Title"
	message.AddExtra("name", "wp123")
	// option 对象，表示推送可选项
	options := push.NewOptions()
	// iOS 平台，是否推送生产环境，false 表示开发环境；如果不指定，就是生产环境
	options.ApnsProduction = true
	// Options 的 Validate 方法会对 time_to_live 属性做范围限制，以满足 JPush 的规范
	options.TimeToLive = 10000000
	// Options 的 Validate 方法会对 big_push_duration 属性做范围限制，以满足 JPush 的规范
	options.BigPushDuration = 1500

	payload := push.NewPushObject()
	payload.Platform = platform
	payload.Audience = audience
	payload.Notification = notification
	payload.Message = message
	//payload.Options = options

	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("json.Marshal PushObject failed:", err)
	}
	fmt.Println("payload:", string(data), "\n")

	// Push 会推送到客户端
	// result, err := client.Push(payload)
	//	showErrOrResult("Push", result, err)

	// PushValidate 的参数和 Push 完全一致
	// 区别在于，PushValidate 只会验证推送调用成功，不会向用户发送任何消息
	result, err := jpushClient.Push(payload)
	fmt.Println("Push", result, err)
}

//CreatePost...
//@Title CreatePost
// @Description CreatePost the Post
// @Param	tagId		query 	string	true		"The id you want to modify"
// @Param	title		query 	string	true		"The id you want to modify"
// @Param	content		query 	string	true		"The id you want to modify"
// @Success 200 {string}成功
// @Failure 403 获取数据失败
//@router /createpost [post]
func (u *UserController) CreatePost() {

	tagId, erro := strconv.Atoi(u.GetString("tagId"))
	if erro != nil {
		responseData.Status = 403
		responseData.Message = erro.Error()
	} else {
		title := u.GetString("title")
		content := u.GetString("content")
		erro = models.CreatePost(tagId, title, content)
		if (erro != nil) {
			responseData.Status = 403
			responseData.Message = erro.Error()
		} else {
			responseData.Status = 200
			responseData.Message = "创建成功"
		}
	}

	u.Data["json"] = responseData.Response
	u.ServeJSON()
}

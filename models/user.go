package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"github.com/astaxie/beego/orm"
	"strconv"
	"crypto/md5"
	"io"
	"math"
	"time"
)

type UserLogin struct {
	Session string `json:"session"`
	User    *User
}

type User struct {
	Id       int    `orm:"column(id);auto",json:"id"`
	Name     string `orm:"column(name);size(10)" json:"name"`
	Password string `orm:"column(password);size(255)"json:"-"` //json:"-"忽略此字段
	Address  string `json:"address",orm:"column(address);size(20);null"`
	Age      int    `json:"age",orm:"column(age);null"`
	Email    string `json:"email",orm:"column(email);size(20);null"`
	Birthday string `json:"birthday",orm:"column(birthday);size(20);null"`
	Phone    string `json:"phone",orm:"column(phone);null;size(11)"`
	//Birthday string `orm:"_"`
	Post *Post `orm:"rel(fk);on_delete(do_nothing);null" json:"-"`
}
type Post struct {
	Id         int `json:"id"`
	Title      string `json:"title"orm:"null"`
	Tags       []*Tag `orm:"rel(m2m);null;rel_table(post_tag_rel)"json:"tags"`
	Content    string `json:"content",orm:"column(content);size(100);null"`
	CreateDate time.Time `orm:"auto_now_add;type(datetime)" json:"create_date"`
	//点赞数量
	PraiseNum int64 `json:"praise_num",orm:"column(praise_num);null"`
	//踩数量
	TreadNum int64 `json:"praise_num",orm:"column(praise_num);null"`
	User     *User `orm:"rel(fk);on_delete(do_nothing);null",json:"user"`
}
type Tag struct {
	Id    int `json:"id"`
	Name  string `orm:"null;" json:"name"`
	Posts []*Post `orm:"reverse(many)" json:"-"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User), new(Post), new(Tag))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int) (*User, error) {
	o := orm.NewOrm()
	//关系查询1
	//v = &User{Id: id}
	//var v1 User
	//qs:=o.QueryTable(new(User))
	//if err := qs.Exclude("address__isnull",true).Filter("id",id).One(&v1); err == nil {
	//	o.Read(v1.Post)
	//	return &v1, nil
	//}else{
	//	return nil, err
	//}

	//关系查询2
	user := &User{}
	if err := o.QueryTable(new(User)).Filter("id", id).RelatedSel().One(user); err == nil {
		fmt.Println(user.Post)
		return user, err
	} else {
		return nil, err
	}

}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...) //排序
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []User
	qs = qs.OrderBy(sortFields...)
	//, fields...
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				if v.Post != nil {
					o.Read(v.Post)
					//读取多对多的数据
					o.QueryTable("tag").Filter("Posts__Post__Id", v.Post.Id).All(&v.Post.Tags)
				}
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func Login(name string, password string) (bool, *User) {
	o := orm.NewOrm()
	values := []string{name, Md5(password)}
	var u *User
	r := o.Raw("SELECT * FROM user WHERE name = ? && password=?", values).QueryRow(&u)
	if r == nil {
		return true, u
	}
	return false, u
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
func RegisterUser(username string, password string) error {
	o := orm.NewOrm()
	u := User{Name: username}
	if r := o.Read(&u, "name"); r == nil {
		return errors.New("该用户名存在，请重新输入用户名")
	} else {
		_, erro := o.Insert(&User{Name: username, Password: Md5(password)})
		if erro != nil {
			return erro
		} else {
			return nil
		}
	}
}

//根据id修改密码
func ModifyPwd(uid string, password string) error {
	o := orm.NewOrm()
	id, erro := strconv.Atoi(uid)
	if erro != nil {
		return erro
	} else {
		v := &User{Id: id}
		erro := o.Read(v)
		if erro != nil {
			return errors.New("该用户不存在")
		} else {
			v.Password = Md5(password)
			_, erro := o.Update(v)
			if erro == nil {
				return nil
			} else {
				return erro
			}
		}
	}
}

/**
获取所有的Tag
 */
func GetAllTag() ([]*Tag, error) {
	o := orm.NewOrm()
	var tags []*Tag
	if _, err := o.QueryTable(new(Tag)).All(&tags); err != nil {
		return nil, err
	} else {
		return tags, err
	}
}

/**
根据tag来获取文章
 */
func GetPostFromTag(tagId int) ([]*Post, error) {
	o := orm.NewOrm()
	var posts []*Post
	_, erro := o.QueryTable(new(Post)).Filter("Tags__Tag__Id", tagId).All(&posts)
	if erro != nil {
		return nil, erro
	}
	return posts, nil
}

/**
根据文章id获取获取文章
 */
func GetPostFromId(postID int) *Post {
	o := orm.NewOrm()
	post := Post{Id: postID}
	if erro := o.Read(&post); erro != nil {
		return nil
	}
	o.Read(post.User)
	return &post
}

/**
获取所有的文章
 */
func GetAllPost(limit int, offset int) ([]*Post, error, int64) {
	o := orm.NewOrm()
	var posts []*Post
	count, erro := o.QueryTable(new(Post)).Limit(limit, offset).Distinct().OrderBy("-create_date").Count()
	if erro != nil {
		return nil, erro, 0
	}
	if _, erro := o.QueryTable(new(Post)).Limit(limit, offset).Distinct().OrderBy("-create_date").All(&posts); erro != nil {
		return nil, erro, 0
	} else {
		for _, post := range posts {
			o.QueryTable(new(Tag)).Filter("Posts__Post__Id", post.Id).One(&post.Tags)
		}
		return posts, erro, count
	}
}

func CreatePost(tagId int, title string, content string) error {
	o := orm.NewOrm()
	tag := &Tag{Id: tagId}
	erro := o.Read(tag)
	if erro == nil {
		post := Post{Title: title, Content: content}
		_, erro := o.Insert(&post)
		if erro == nil {
			m2m := o.QueryM2M(&post, "Tags")
			_, erro := m2m.Add(tag)
			if erro == nil {
				return erro
			}
			return erro
		}
		return erro
	}
	return erro
}

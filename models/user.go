package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"github.com/astaxie/beego/orm"
	"strconv"
	"crypto/md5"
	"encoding/hex"
)

type UserLogin struct {
	Session string `json:"session"`
	User    *User
}

type User struct {
	Id       int    `orm:"column(id);auto"json:"id"`
	Name     string `orm:"column(name);size(10)" json:"name"`
	Password string `orm:"column(password);size(255)"json:"-"` //json:"-"忽略此字段
	Address  string `json:"address"orm:"column(address);size(20);null"`
	Age      int    `json:"age"orm:"column(age);null"`
	Email    string `json:"email"orm:"column(email);size(20);null"`
	Birthday string `json:"birthday"orm:"column(birthday);size(20);null"`
	//Birthday string `orm:"_"`
	Post *Post `orm:"rel(fk);on_delete(do_nothing);null" json:"post"`
}
type Post struct {
	Id    int `json:"id"`
	Title string `json:"title",orm:"null"`
	Tags  []*Tag `orm:"rel(m2m);null;rel_table(post_tag_rel)",json:"tags"`
}

//type PostTagRel struct{
//	Id int `json:"id"`
//	Posts *Post `json:"posts"`
//	Tags *Tag `json:"tags"`
//}

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

//md5加密        return 加密 后的字符串
//values  待加密的字符串
func Md5(values string) string {
	h := md5.New()
	h.Write([]byte(values + "wp"))
	return hex.EncodeToString(h.Sum(nil))
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

//func GAO()  {
//	o:=orm.NewOrm()
//	qs:=o.QueryTable("user")
//	qs.Filter("name","wp123")
//
//
//	qs.Exclude("address__isnull",true).Filter("name","wp1234")
//
//
//	qs.Limit(10)//限制10条数据
//
//	qs.Limit(10,20)
//
//	qs.Limit(-1)//不限制
//
//	qs.GroupBy("id","age")
//
//	qs.OrderBy("id")
//	qs.Distinct()//返回 不重复的数据
//	//age在原来的基础上增加10    支持加减乘除
//	qs.Update(orm.Params{"name":"wp","age":orm.ColValue(orm.ColAdd,10),"address":"gs供电所"})
//
//
//	i,_:=qs.PrepareInsert()
//	i.Insert(User{})
//}

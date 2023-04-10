package complex

import "time"

// 项目目录
type Project struct {
	Id    int64  `form:"-"`
	Code  string `orm:"null"`                                              //编号
	Title string `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
}

// 目录里的成果表
type Product struct {
	Id           int64
	Code         string `orm:"null"` //编号
	Title        string `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"`
	Label        string `orm:"null"` //标签
	Uid          int64  `orm:"null"`
	Principal    string `orm:"null"`       //提供人
	ProjectId    int64  `orm:"null"`       //目录projectid
	TopProjectId int64  `orm:"default(0)"` //项目id
}

// 用户表
type User struct {
	Id       int64
	Username string `json:"name",orm:"unique"` //这个拼音的简写
	Nickname string //中文名
}

// 文章表，文章放在成果下面，所以文章没有直接对应作者，而是通过product成果来查作者
type Article struct {
	Id        int64     `json:"id",form:"-"`
	Subtext   string    `orm:"sie(20)"`
	Content   string    `json:"html",orm:"sie(5000)"`
	ProductId int64     `orm:"null"` //成果下面的文章
	Views     int64     `orm:"default(0)"`
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now_add;type(datetime)"`
}

// 查询返回新建的结构体-取出用户文章数目
type Result struct {
	Usernickname string `json:"name"`
	Productid    int64
	Total        int64 `json:"value"`
}

func GetWxUserArticles(pid int64) (results []*Result, err error) {
	db := GetDB()
	db.Order("total desc").Table("article").Select("product_id as productid, count(*) as total,user.nickname as usernickname").
		Joins("left JOIN product on product.id = article.product_id").
		Joins("left JOIN user on user.id = product.uid").Group("product.uid").
		Joins("left JOIN project on project.id = product.project_id").Where("project.id=?", pid).
		Scan(&results)
	return results, err
}

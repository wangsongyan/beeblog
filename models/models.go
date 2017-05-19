package models

import (
	"os"
	"path"
	"strconv"
	"time"

	"github.com/UnKnwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"auto_now_add;type(datetime);index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index;null"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Category        string
	Created         time.Time `orm:"auto_now_add;type(datetime);index;"`
	Updated         time.Time `orm:"auto_now;type(datetime);index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index;null"`
	ReplyCount      int64
	ReplyLastUserId int64
}

type Comment struct {
	Id       int64
	Tid      int64
	Nickname string
	Content  string    `orm:"size(1000)"`
	Created  time.Time `orm:"index"`
}

func RegisterDB() {

	orm.Debug = true
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)

}

func AddCategory(name string) error {

	o := orm.NewOrm()
	cate := &Category{Title: name, Created: time.Now(), TopicTime: time.Now()}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func GetAllCategories() ([]*Category, error) {

	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

// 文章
func AddTopic(title, category, content string) error {

	o := orm.NewOrm()

	topic := &Topic{
		Title:    title,
		Category: category,
		Content:  content,
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	_, err := o.Insert(topic)
	if err != nil {
		return err
	}

	// 分类文章数量修改
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		cate.TopicCount++
		_, err = o.Update(cate)
	}
	return err
}

func GetAllTopics(isDesc bool, cate string) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error
	if isDesc {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err
}

func CountTopicByCategory(cate string) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("topic")
	count, err := qs.Filter("category", cate).Count()
	return count, err
}
func GetTopic(id string) (*Topic, error) {

	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tid).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}

func ModifyTopic(id, title, category, content string) error {

	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{
		Id: tid,
	}

	if err = o.Read(topic); err == nil {
		topic.Title = title
		topic.Category = category
		topic.Content = content
		topic.Updated = time.Now()
		if _, err = o.Update(topic); err != nil {
			return err
		}
		return nil
	}
	return err

}

func DeleteTopic(id string) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{
		Id: tid,
	}
	var cate string
	if err = o.Read(topic); err == nil {
		cate = topic.Category
	}
	_, err = o.Delete(topic)
	if err != nil {
		return err
	}
	category := new(Category)
	err = o.QueryTable("category").Filter("title", cate).One(category)
	if err == nil {
		if category.TopicCount > 1 {
			category.TopicCount--
			_, err = o.Update(category)
		}
	}
	return err
}

//
func AddReply(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	o := orm.NewOrm()
	comment := &Comment{
		Tid:      tidNum,
		Nickname: nickname,
		Content:  content,
		Created:  time.Now(),
	}
	_, err = o.Insert(comment)
	return err
}

//
func GetRepliesByTopicId(tid string) ([]*Comment, error) {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	comments := make([]*Comment, 0)
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", id).OrderBy("-created").All(&comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

//
func DeleteReply(id string) error {
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	comment := &Comment{
		Id: idNum,
	}
	o := orm.NewOrm()
	_, err = o.Delete(comment)
	return err
}

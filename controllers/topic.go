package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	"path"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.TplName = "topic.html"
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	topics, err := models.GetAllTopics(false, "", "")
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}
}

func (this *TopicController) Post() {

	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	title := this.GetString("title")
	content := this.GetString("content")
	category := this.GetString("category")
	label := this.GetString("label")
	id := this.GetString("id")

	var attachment string
	_, fileheader, err := this.GetFile("attachment")
	if err == nil {
		attachment = fileheader.Filename
		// 拷贝文件到指定目录
		beego.Info(attachment)
		err = this.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			beego.Error(err)
		}
	} else {
		beego.Error(err)
	}

	if len(id) == 0 {
		err = models.AddTopic(title, category, label, content, attachment)
	} else {
		err = models.ModifyTopic(id, title, category, label, content, attachment)
	}

	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)

}

func (this *TopicController) Add() {
	this.TplName = "topic_add.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
}

func (this *TopicController) View() {
	this.TplName = "topic_view.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	id := this.Ctx.Input.Params()["0"]
	topic, err := models.GetTopic(id)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	} else {
		this.Data["Topic"] = topic
		this.Data["Tid"] = id
		this.Data["Labels"] = strings.Split(topic.Label, " ")
	}

	replies, err := models.GetRepliesByTopicId(id)
	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["Replies"] = replies
}

func (this *TopicController) Modify() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	this.TplName = "topic_modify.html"
	id := this.Ctx.Input.Params()["0"]
	topic, err := models.GetTopic(id)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	} else {
		this.Data["Topic"] = topic
		this.Data["Tid"] = id
	}
}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	id := this.Ctx.Input.Params()["0"]
	if len(id) > 0 {
		err := models.DeleteTopic(id)
		if err != nil {
			beego.Error(err)
		}
	}
	this.Redirect("/", 302)
	return
}

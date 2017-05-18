package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.TplName = "topic.html"
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	topics, err := models.GetAllTopics(false, "")
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
	id := this.GetString("id")

	var err error
	if len(id) == 0 {
		err = models.AddTopic(title, category, content)
	} else {
		err = models.ModifyTopic(id, title, category, content)
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

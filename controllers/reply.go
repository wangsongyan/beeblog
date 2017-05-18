package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Add() {
	tid := this.GetString("tid")
	nickname := this.GetString("nickname")
	content := this.GetString("content")
	err := models.AddReply(tid, nickname, content)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic/view/"+tid, 302)
}

func (this *ReplyController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	tid := this.Ctx.Input.Params()["0"]
	id := this.Ctx.Input.Params()["1"]
	if len(id) > 0 {
		err := models.DeleteReply(id)
		if err != nil {
			beego.Error(err)
		}
	}
	this.Redirect("/topic/view/"+tid, 302)
}

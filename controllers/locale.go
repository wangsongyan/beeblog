package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type baseContoller struct {
	beego.Controller
	i18n.Locale
}

func (this *baseContoller) Prepare() {

	lang := this.GetString("lang")
	if lang == "zh-CN" {
		this.Lang = lang
	} else {
		this.Lang = "en-US"
	}
	this.Data["Lang"] = this.Lang

}

type LocaleController struct {
	baseContoller
}

func (this *LocaleController) Get() {
	this.TplName = "locale.html"
	// this.Data["hi"] = i18n.Tr(this.Lang, "hi")
	// this.Data["hey"] = i18n.Tr(this.Lang, "hey")
	this.Data["hi"] = "hi"
	this.Data["hey"] = "hey"
	this.Data["about"] = "about"
}

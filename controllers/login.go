package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {

	isExit := this.GetString("exit") == "true"
	if isExit {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 302)
		return
	}

	this.TplName = "login.html"
}

func (this *LoginController) Post() {

	//this.Ctx.WriteString(fmt.Sprint(this.Ctx.Input.Params()))

	uname := this.GetString("uname")
	pwd := this.GetString("pwd")
	autoLogin := this.GetString("autoLogin") == "on"

	fmt.Println(fmt.Sprint(this.Input()))

	if uname == beego.AppConfig.String("uname") {
		if pwd == beego.AppConfig.String("pwd") {
			maxAge := 0
			if autoLogin {
				maxAge = 1<<32 - 1
			}
			this.Ctx.SetCookie("uname", uname, maxAge, "/")
			this.Ctx.SetCookie("pwd", pwd, maxAge, "/")

		} else {

		}
	}

	this.Redirect("/", 302)
	return

}

func checkAccount(ctx *context.Context) bool {

	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value
	return uname == beego.AppConfig.String("uname") && pwd == beego.AppConfig.String("pwd")
}

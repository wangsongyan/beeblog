package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {

	op := this.GetString("op")
	switch op {
	case "add":
		name := this.GetString("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 302)
		return
	case "del":
		id := this.GetString("id")
		if len(id) == 0 {
			break
		}
		err := models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 302)
		return
	}

	this.Data["IsCategory"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "category.html"

	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}

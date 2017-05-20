package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	cate := c.GetString("cate")
	label := c.GetString("label")
	topics, err := models.GetAllTopics(true, cate, label)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Topics"] = topics
	}
	c.TplName = "home.html"

	categories, err := models.GetAllCategories()
	// for _, category := range categories {
	// 	count, err := models.CountTopicByCategory(category.Title)
	// 	if err == nil {
	// 		category.TopicCount = count
	// 	}
	// }

	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Categories"] = categories
	}
}

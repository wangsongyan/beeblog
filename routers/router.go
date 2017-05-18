package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/test", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.AutoRouter(&controllers.ReplyController{})
}

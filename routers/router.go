package routers

import (
	"github.com/astaxie/beego"
	"github.com/wangsongyan/beeblog/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
}

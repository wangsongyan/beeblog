package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
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

	// 附件处理方式一
	//beego.SetStaticPath("/attachment", "attachment")
	// 附件处理方式二
	beego.Router("/attachment/:all", &controllers.FileController{})

	// 国际化支持
	i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini")
	i18n.SetMessage("en-US", "conf/locale_en-US.ini")
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Router("/locale", &controllers.LocaleController{})

}

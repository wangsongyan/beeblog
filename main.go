package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/wangsongyan/beeblog/models"
	_ "github.com/wangsongyan/beeblog/routers"
)

func init() {
	models.RegisterDB()
}

func main() {

	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Run()
}

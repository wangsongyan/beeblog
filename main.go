package main

import (
	"beeblog/models"
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	models.RegisterDB()
}

func main() {

	orm.Debug = true
	//orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Shanghai") //time.UTC //time.Local
	orm.DefaultTimeLoc = time.UTC
	orm.RunSyncdb("default", false, true)
	beego.Run()
}

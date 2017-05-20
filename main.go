package main

import (
	"beeblog/models"
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
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
	os.Mkdir("attachment", os.ModePerm)
	beego.Run()
}

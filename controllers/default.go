package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type Entity struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
	Index int    `json:"index"`
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	c.Data["TrueCond"] = true
	c.Data["FalseCond"] = false
	c.TplName = "index.tpl"

	type u struct {
		Name   string
		Age    int
		Gender string
	}

	user := &u{
		Name:   "wangsongyan",
		Age:    25,
		Gender: "male",
	}

	c.Data["user"] = user
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	c.Data["nums"] = a

	c.Data["TplVar"] = "hey guys"
	c.Data["Html"] = "<div>Hello beego</div>"
	c.Data["Pipe"] = "<div>Hello beego</div>"

	cates := make([]Entity, 0)
	for i := 0; i < 10; i++ {
		cate := Entity{
			Title: "test",
			Index: i,
		}
		cates = append(cates, cate)
	}
	fmt.Println(cates)
	jsonstr, err := json.Marshal(cates)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonstr))
	c.Data["Categoris"] = string(jsonstr)

}

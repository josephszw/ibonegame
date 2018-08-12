package routers

import (
	"ibgamemanage/controllers"

	"github.com/astaxie/beego"
)

func init() {

	//查看博客详细信息
	beego.Router("/view/:PlayerId([0-9]+)", &controllers.ViewController{})
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.LoginController{})
	//新建博客博文
	beego.Router("/new", &controllers.NewController{})
	//删除博文
	beego.Router("/delete/:id([0-9]+)", &controllers.DeleteController{})
	//编辑博文
	beego.Router("/edit/:id([0-9]+)", &controllers.EditController{})
}

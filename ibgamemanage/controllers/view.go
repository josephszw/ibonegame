package controllers

import (
	"ibgamemanage/models"

	"github.com/astaxie/beego"
	"github.com/spf13/cast"
)

type ViewController struct {
	beego.Controller
}

func (this *ViewController) Get() {
	id := cast.ToInt64(this.Ctx.Input.Params()[":PlayerId"])
	this.Data["info"] = models.GetPlayerInfo(id)
	if user, ok := this.GetSecureCookie("mycookie", "token"); ok {
		this.Data["user"] = user
	}
	this.TplName = "billView.html"
}

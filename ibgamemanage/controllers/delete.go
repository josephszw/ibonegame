package controllers

import (
	"ibgamemanage/models"

	"github.com/spf13/cast"

	"github.com/astaxie/beego"
)

type DeleteController struct {
	beego.Controller
}

func (this *DeleteController) Get() {
	id := cast.ToInt64(this.Ctx.Input.Params()[":id"])
	pi := models.GetPlayerInfo(id)
	this.Data["Post"] = pi
	models.DelPlayerInfo(pi)
	this.Ctx.Redirect(302, "/")
}

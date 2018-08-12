package controllers

import (
	"ibgamemanage/models"

	"github.com/astaxie/beego"
	"github.com/spf13/cast"
)

type NewController struct {
	beego.Controller
}

func (this *NewController) Get() {
	this.TplName = "billAdd.html"
}

func (this *NewController) Post() {
	inputs := this.Input()
	var playerInfo models.PlayerInfo
	playerInfo.Name = inputs.Get("name")
	playerInfo.Nick_name = inputs.Get("nickname")
	playerInfo.Position = cast.ToInt(inputs.Get("postion"))
	playerInfo.Second_position = cast.ToInt(inputs.Get("second"))
	playerInfo.Type = cast.ToInt(inputs.Get("type"))
	playerInfo.Score = cast.ToInt(inputs.Get("score"))
	playerInfo.Rebound = cast.ToInt(inputs.Get("rebound"))
	playerInfo.Assist = cast.ToInt(inputs.Get("assist"))
	playerInfo.Steal = cast.ToInt(inputs.Get("steal"))
	playerInfo.Cap = cast.ToInt(inputs.Get("cap"))
	playerInfo.Appear_num = cast.ToInt(inputs.Get("num"))
	models.SavePlayerInfo(playerInfo)
	this.Ctx.Redirect(302, "/")
}

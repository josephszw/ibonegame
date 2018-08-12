package controllers

import (
	"ibgamemanage/models"

	"github.com/astaxie/beego"
	"github.com/beego/wetalk/modules/utils"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {

	count := models.GetPlayerInfoCount()
	if count > 0 {
		pagesize := 12
		p := utils.NewPaginator(this.Ctx.Request, pagesize, count)
		this.Data["paginator"] = p
		this.Data["infos"] = models.GetPlayerInfoLimit(p.Offset(), pagesize)
	}
	if user, ok := this.GetSecureCookie("mycookie", "token"); ok {
		this.Data["user"] = user
	}
	this.TplName = "billList.html"
}

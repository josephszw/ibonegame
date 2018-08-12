package controllers

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"ibgame/models/user_model"
	"io"
	"regexp"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.tpl"
	// this.Ctx.Request.ParseForm()
	// username := this.Ctx.Request.Form.Get("username")
	// password := this.Ctx.Request.Form.Get("password")
	// md5Password := md5.New()
	// io.WriteString(md5Password, password)
	// buffer := bytes.NewBuffer(nil)
	// fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	// userInfo, _ := user_model.Get(username, password)
	// if userInfo[username] == "success" {
	// 	//登录成功设置session
	// 	expiration := time.Now()
	// 	expiration = expiration.AddDate(1, 0, 0)
	// 	this.SetSecureCookie("mycookie", "token", username)
	// 	this.Ctx.Redirect(302, "/")
	// }
}
func (this *LoginController) Post() {
	inputs := this.Input()
	username := inputs.Get("username")
	password := inputs.Get("password")
	md5Password := md5.New()
	io.WriteString(md5Password, password)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	userInfo, _ := user_model.Get(username, password)
	if userInfo[username] == "success" {
		//登录成功设置session
		expiration := time.Now()
		expiration = expiration.AddDate(1, 0, 0)
		this.SetSecureCookie("mycookie", "token", username, expiration)
		this.Ctx.Redirect(302, "/")
	} else {
		this.Ctx.Redirect(302, "/login")
	}
}

func checkPassword(password string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password); !ok {
		return false
	}
	return true
}

func checkUsername(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", username); !ok {
		return false
	}
	return true
}

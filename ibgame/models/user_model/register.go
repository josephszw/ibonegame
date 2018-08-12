package user_model

import (
	"encoding/base64"
	"errors"
	"fmt"
	"ibgame/logs"
	"ibgame/models/mysql"
	"ibgame/utils"
	"math/rand"
	"strings"
	"time"

	"github.com/spf13/cast"

	"github.com/go-xorm/xorm"
)

const (
	reKey = "amlhbWlzdXpoYW53ZWlxdw=="
)

//用户信息
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Register 注册
func Register(param RegisterParam) (ret map[string]string, err error) {
	var engine *xorm.Engine
	ret = make(map[string]string)
	var user UserLogin
	if engine, err = mysql.GetEngine(); err != nil {
		logs.Error.Println("engine error", err)
		return
	}
	pass := utils.Checkmail(param.Mail)
	if !pass {
		err = errors.New("邮箱无效")
		return
	}
	if len(param.Token) <= 0 {
		err = errors.New("token 无效")
		return
	}
	str := utils.Aes128CBCDecrypt(param.Token, reKey, false)
	logs.Info.Println(str)
	now := time.Now().Unix()
	user.Name = param.Name
	user.Password = str
	uid := randSeq(10)
	user.Uid = uid
	user.Cid = param.Cid
	user.Ctime = now
	user.Mail = param.Mail
	user.MailVerify = 0
	user.Utime = now
	if _, e := engine.Insert(&user); e != nil {
		logs.Error.Println("insert error", e)
		err = e
		return
	}
	at := generateAuthToken(uid, now)
	ret[param.Name] = fmt.Sprintf("uid是%s http://localhost:12356/parseauthtoken?token=%s", uid, at)
	var s utils.SendMailParam
	s.Content = fmt.Sprintf("http://localhost:12356/parseauthtoken?token=%s", at)
	s.Title = "请鉴权"
	s.ToMail = param.Mail
	s.ToName = param.Name
	err = utils.SendMail(s)
	if err != nil {
		logs.Error.Println(err)
		return
	}
	return
}

var letters = []rune("1234567890")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// GenerateAuthToken 生成鉴权
func generateAuthToken(uid string, time int64) (ret string) {

	str := fmt.Sprintf("%s/%s", uid, cast.ToString(time))
	data := base64.StdEncoding.EncodeToString([]byte(str))
	ret = utils.Aes128CBCEncrypt(data, reKey)
	return
}

//ParseAuthToken 解析
func ParseAuthToken(str string) (err error) {

	var engine *xorm.Engine
	users := make([]UserLogin, 0)
	var user UserLogin
	if engine, err = mysql.GetEngine(); err != nil {
		logs.Error.Println("engine error", err)
	}

	data := utils.Aes128CBCDecrypt(str, reKey, false)

	b, _ := base64.StdEncoding.DecodeString(data)

	r := strings.Split(string(b), "/")
	uid := r[0]
	time := cast.ToInt64(r[1])

	if e := engine.Where("uid=? and ctime=?", uid, time).Find(&users); e != nil {
		logs.Error.Println("db error", err)
		return
	} else {
		if len(users) <= 0 {
			err = errors.New("鉴权失败")
			return
		} else {
			user = users[0]
			user.MailVerify = 1
			if _, e1 := engine.Cols("mail_verify").Update(&user); e1 != nil {
				err = errors.New("鉴权失败")
				return
			}
		}
	}
	return
}

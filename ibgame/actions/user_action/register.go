package user_action

import (
	"encoding/json"
	"ibgame/logs"
	"ibgame/models/user_model"
	"io/ioutil"
	"net/http"
	"strings"
)

// Register 注册action
func Register(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logs.Error.Println("ioutil.ReadAll err", err)
		return
	}
	logs.Info.Println("ServeHTTP", "reqeust = ", string(data))
	var param user_model.RegisterParam
	err = json.Unmarshal(data, &param)

	if param.Name == "" {
		re := "注册名字为空"
		ret := map[string]interface{}{"code": 401, "msg": "error", "data": re}
		bytes, err := json.MarshalIndent(ret, " ", "    ")
		if err != nil {
			logs.Error.Println("json:err", err)
		}
		w.Write(bytes)
		return
	}
	if re, e := user_model.Register(param); e != nil {
		logs.Error.Println("models.Register:err", err)
	} else {
		logs.Info.Println(re)
		ret := map[string]interface{}{"code": 0, "msg": "ok", "data": re}
		bytes, err := json.MarshalIndent(ret, " ", "    ")
		if err != nil {
			logs.Error.Println("json:err", err)
		}
		w.Write(bytes)
	}
}

// ParseAuthToken token解析
func ParseAuthToken(w http.ResponseWriter, r *http.Request) {
	t := r.FormValue("token")
	str := strings.Replace(t, " ", "+", -1)
	logs.Info.Println("t", str)
	user_model.ParseAuthToken(str)
}

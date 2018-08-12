package user_action

import (
	"encoding/json"
	"ibgame/logs"
	"ibgame/models/user_model"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	uname := r.FormValue("username")
	pword := r.FormValue("password")

	if re, e := user_model.Get(uname, pword); e != nil {
		logs.Error.Println("models.Get:err", e)
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

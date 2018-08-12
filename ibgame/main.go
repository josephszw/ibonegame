package main

import (
	"fmt"
	"ibgame/actions/getui_action"
	"ibgame/actions/user_action"
	"ibgame/logs"
	"net/http"
)

func main() {
	http.HandleFunc("/login", user_action.Login)
	http.HandleFunc("/register", user_action.Register)
	http.HandleFunc("/pushsingle", getui_action.PushSingle)
	http.HandleFunc("/parseauthtoken", user_action.ParseAuthToken)
	err := http.ListenAndServe(":12356", nil) //设置监听的端口
	if err != nil {
		logs.Error.Println("ListenAndServe:error")
	} else {
		fmt.Println("监听12356")
	}
}

package getui_model

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"ibgame/logs"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func PushSingel(param Single) (ret string, err error) {
	if len(param.Cid) == 0 {
		return "", fmt.Errorf("[PushToSingle] 错误的目标设备, cid 与 alias 任选且必选一个")
	}
	param.Requestid = strconv.FormatInt(time.Now().UnixNano(), 12)
	// 构造请求
	data, _ := json.Marshal(param)
	req, err := http.NewRequest("POST", "https://restapi.getui.com/v1/"+APPID+"/push_single", ioutil.NopCloser(bytes.NewReader(data)))
	if err != nil {
		return "", fmt.Errorf("[PushToSingle] 创建 发送单客户端信息 请求失败, err: %s", err)
	}
	var a Auth
	ts := fmt.Sprintf("%d", int64(time.Now().UnixNano()/1000000))
	sign := sha256.Sum256([]byte(APPKEY + ts + MASTERSECRET))
	signStr := fmt.Sprintf("%x", sign)
	a.Sign = signStr
	a.Timestamp = ts
	a.Appkey = APPKEY
	token, _ := getAuth(a)
	req.Header["Content-Type"] = []string{"application/json"}
	req.Header["authtoken"] = []string{token}
	// 发送请求
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("[PushToSingle] 发送 单客户端信息 请求失败, err: %s", err)
	}
	defer rsp.Body.Close()
	// 解析-body
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", fmt.Errorf("[PushToSingle] 发送 单客户端信息请求 返回的body无法解析, err: %s", err)
	}
	var re RspBody
	err = json.Unmarshal(rspBody, &re)
	if err != nil {
		return "", fmt.Errorf("[PushToSingle] 发送 单客户端信息 请求返回的JSON无法解析, err: %s", err)
	}
	if re.Result != "ok" {
		return "", fmt.Errorf("[PushToSingle] 发送 单客户端信息 请求不成功, ret: %v", ret)
	}
	ret = string(rspBody)
	return
}

func getAuth(p Auth) (ret string, err error) {
	payloadBytes, e := json.Marshal(p)
	if e != nil {
		logs.Error.Println("http resp", e)
		err = e
		return
	}
	body := bytes.NewReader(payloadBytes)
	req, e1 := http.NewRequest("POST", "https://restapi.getui.com/v1/"+APPID+"/auth_sign", body)
	if e1 != nil {
		logs.Error.Println("http resp", e1)
		err = e1
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, e2 := http.DefaultClient.Do(req)
	if e2 != nil {
		logs.Error.Println("http resp", e1)
		err = e2
		return
	}
	b, e3 := ioutil.ReadAll(resp.Body)
	if e3 != nil {
		logs.Error.Println("http resp", e3)
		err = e3
		return
	}
	var re Result
	_ = json.Unmarshal(b, &re)
	ret = re.AuthToken
	defer resp.Body.Close()
	return
}

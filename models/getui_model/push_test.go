package getui_model

import (
	"fmt"
	"testing"
)

func TestSingle(t *testing.T) {

	var p Single

	p.Cid = "2326fbc31ef63427a82e667b74353341"
	p.Message.Appkey = "al0zZ6nvSO9tvxPPrTVHD9"
	p.Message.IsOffline = false
	p.Message.Msgtype = "notification"
	p.Message.OfflineExpireTime = 1000000
	p.Notification.Style.Type = 0
	p.Notification.Style.Title = "Title"
	p.Notification.Style.Text = "Text"
	p.Notification.TransmissionContent = ""
	p.Notification.TransmissionType = false
	p.Requestid = "123456"

	re, e := PushSingel(p)
	fmt.Println("re", re)
	fmt.Println("e", e)
}

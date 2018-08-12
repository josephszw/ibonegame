package utils

import "github.com/go-gomail/gomail"

// SendMailParam 邮件参数
type SendMailParam struct {
	ToMail  string
	ToName  string
	Title   string
	Content string
}

// SendMail 发邮件
func SendMail(p SendMailParam) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", "545397649@qq.com") // 发件人
	m.SetHeader("To",                       // 收件人
		m.FormatAddress(p.ToMail, p.ToName),
	)
	m.SetHeader("Subject", p.Title)   // 主题
	m.SetBody("text/html", p.Content) // 正文
	d := gomail.NewDialer("smtp.qq.com", 465, "545397649@qq.com", "ookdjcmikjorbgaj")
	if e := d.DialAndSend(m); e != nil {
		err = e
		panic(e)
	}
	return
}

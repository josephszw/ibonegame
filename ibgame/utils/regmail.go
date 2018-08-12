package utils

import (
	"regexp"
)

const (
	reg = "^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
)

// Checkmail 邮箱正则
func Checkmail(email string) bool {
	regMail := regexp.MustCompile(reg)
	return regMail.MatchString(email)
}

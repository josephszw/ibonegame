package strucct

type UserLogin struct {
	Id         int    `xorm:"not null pk autoincr INT(10)"`
	Name       string `xorm:"not null default '' index VARCHAR(100)"`
	Uid        string `xorm:"not null default '' unique VARCHAR(40)"`
	Password   string `xorm:"not null default '' VARCHAR(40)"`
	Mail       string `xorm:"not null default '' VARCHAR(40)"`
	Ctime      int64  `xorm:"not null default 0 BIGINT(20)"`
	Cid        string `xorm:"not null default '' VARCHAR(40)"`
	MailVerify int    `xorm:"not null default 0 TINYINT(4)"`
	Utime      int64  `xorm:"not null default 0 BIGINT(20)"`
}

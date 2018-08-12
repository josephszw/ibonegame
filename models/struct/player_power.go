package strucct

type PlayerPower struct {
	Id               int64 `xorm:"pk autoincr BIGINT(20)"`
	PlayerId         int64 `xorm:"not null default 0 index BIGINT(20)"`
	InsideAttack     int   `xorm:"not null default 0 INT(11)"`
	OutsideAttack    int   `xorm:"not null default 0 INT(11)"`
	OffensiveRebound int   `xorm:"not null default 0 INT(11)"`
	DefensiveRebound int   `xorm:"not null default 0 INT(11)"`
	Pass             int   `xorm:"not null default 0 INT(11)"`
	InsideDefense    int   `xorm:"not null default 0 INT(11)"`
	OutsideDefense   int   `xorm:"not null default 0 INT(11)"`
	BasePower        int   `xorm:"not null default 0 INT(11)"`
}

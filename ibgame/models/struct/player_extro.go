package strucct

type PlayerExtro struct {
	Id               int64 `xorm:"not null pk autoincr BIGINT(20)"`
	PlayerId         int64 `xorm:"not null default 0 index BIGINT(20)"`
	OffensiveRebound int   `xorm:"not null default 0 INT(11)"`
	DefensiveRebound int   `xorm:"not null default 0 INT(11)"`
	ShotAttempt      int   `xorm:"not null default 0 INT(11)"`
	ShotMade         int   `xorm:"not null default 0 INT(11)"`
	ThreeAttempt     int   `xorm:"not null default 0 INT(11)"`
	ThreeMade        int   `xorm:"not null default 0 INT(11)"`
	InsideAttempt    int   `xorm:"not null default 0 INT(11)"`
	InsideMade       int   `xorm:"not null default 0 INT(11)"`
	TurnOff          int   `xorm:"not null default 0 INT(11)"`
	Faul             int   `xorm:"not null default 0 INT(11)"`
}

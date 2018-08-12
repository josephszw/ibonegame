package manage_model

//PlayerInfo base info
type PlayerInfo struct {
	PlayerId       int64  `xorm:"not null pk autoincr BIGINT(20)"`
	Name           string `xorm:"not null default '' unique(uk_name) VARCHAR(128)"`
	NickName       string `xorm:"not null default '' unique(uk_name) VARCHAR(128)"`
	Position       int    `xorm:"not null unique(uk_type_position) INT(11)"`
	SecondPosition int    `xorm:"not null default 0 INT(11)"`
	Type           int    `xorm:"not null default 0 unique(uk_type_position) INT(11)"`
	Score          int    `json:""`
	Rebound        int    `json:""`
	Assist         int    `json:""`
	Steal          int    `json:""`
	Cap            int    `json:""`
	AppearNum      int    `json:""`
}

type PlayerPower struct {
	Id               int64 `xorm:"pk autoincr BIGINT(20)"`
	PlayerId         int64 `xorm:"not null default 0 index BIGINT(20)"`
	InsideAttack     int   `json:""`
	OutsideAttack    int   `json:""`
	OffensiveRebound int   `json:""`
	DefensiveRebound int   `json:""`
	Pass             int   `json:""`
	InsideDefense    int   `json:""`
	OutsideDefense   int   `json:""`
	BasePower        int   `json:""`
}

//AddParam 参数
type AddParam struct {
	Name                  string `json:"name"`
	NickName              string `json:"nick_name"`
	Position              int    `json:"position"`
	SecondPosition        int    `json:"second_position"`
	Type                  int    `json:"type"`
	Score                 int    `json:"score"`
	Rebound               int    `json:"rebound"`
	Assist                int    `json:"assist"`
	Steal                 int    `json:"steal"`
	Cap                   int    `json:"cap"`
	AppearNum             int    `json:"appear_num"`
	InsideAttack          int    `json:"inside_attack"`
	OutsideAttack         int    `json:"outside_attack"`
	OffensiveReboundPower int    `json:"offensive_rebound_power"`
	DefensiveReboundPower int    `json:"defensive_rebound_power"`
	Pass                  int    `json:"pass"`
	InsideDefense         int    `json:"inside_defense"`
	OutsideDefense        int    `json:"outside_defense"`
	BasePower             int    `json:"base_power"`
	OffensiveRebound      int    `json:"offensive_rebound"`
	DefensiveRebound      int    `json:"defensive_rebound"`
	ShotAttempt           int    `json:"shot_attempt"`
	ShotMade              int    `json:"shot_made"`
	ThreeAttempt          int    `json:"three_attempt"`
	ThreeMade             int    `json:"three_made"`
	InsideAttempt         int    `json:"inside_attempt"`
	InsideMade            int    `json:"inside_made"`
	TurnOff               int    `json:"turn_off"`
	Faul                  int    `json:"faul"`
}

//PlayerResult 参数
type PlayerResult struct {
	PlayerId              int64  `json:"player_id"`
	Name                  string `json:"name"`
	NickName              string `json:"nick_name"`
	Position              string `json:"position"`
	SecondPosition        string `json:"second_position"`
	Type                  string `json:"type"`
	Score                 int    `json:"score"`
	Rebound               int    `json:"rebound"`
	Assist                int    `json:"assist"`
	Steal                 int    `json:"steal"`
	Cap                   int    `json:"cap"`
	AppearNum             int    `json:"appear_num"`
	InsideAttack          int    `json:"inside_attack"`
	OutsideAttack         int    `json:"outside_attack"`
	OffensiveReboundPower int    `json:"offensive_rebound_power"`
	DefensiveReboundPower int    `json:"defensive_rebound_power"`
	Pass                  int    `json:"pass"`
	InsideDefense         int    `json:"inside_defense"`
	OutsideDefense        int    `json:"outside_defense"`
	BasePower             int    `json:"base_power"`
	OffensiveRebound      int    `json:"offensive_rebound"`
	DefensiveRebound      int    `json:"defensive_rebound"`
	ShotAttempt           int    `json:"shot_attempt"`
	ShotMade              int    `json:"shot_made"`
	ThreeAttempt          int    `json:"three_attempt"`
	ThreeMade             int    `json:"three_made"`
	InsideAttempt         int    `json:"inside_attempt"`
	InsideMade            int    `json:"inside_made"`
	TurnOff               int    `json:"turn_off"`
	Faul                  int    `json:"faul"`
}

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

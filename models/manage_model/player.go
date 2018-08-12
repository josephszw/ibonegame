package manage_model

import (
	"ibgame/logs"
	"ibgame/models/mysql"
)

const (
	superstar = 1  //"当家球星"
	allstar   = 11 //"全明星"
	scorer    = 2  //"得分手"
	defender  = 3  //"防守者"
	threer    = 4  //"三分手"
	maker     = 5  //"组织者"
	rebound   = 6  //"篮板手"
	sixer     = 7  //"第六人"
	threeD    = 8  //"3d"
	tibu      = 9  //"替补"
	banch     = 10 //"板凳"
)

var typeMap = map[int]string{1: "当家球星", 2: "得分手", 3: "防守者", 4: "三分手", 5: "组织者", 6: "篮板手", 7: "第六人", 8: "3d", 9: "替补", 10: "板凳"}
var posMap = map[int]string{1: "控球后卫", 2: "得分后卫", 3: "小前锋", 4: "大前锋", 5: "中锋"}
var posEMap = map[int]string{1: "Pg", 2: "Sg", 3: "SF", 4: "PF", 5: "C"}

// Add 添加球员
func Add(p AddParam) (ret string, err error) {

	var pi PlayerInfo
	var pe PlayerExtro
	var pw PlayerPower
	pi.Name = p.Name
	pi.NickName = p.NickName
	pi.Position = p.Position
	pi.SecondPosition = p.SecondPosition
	pi.Type = p.Type
	pi.Score = p.Score
	pi.Rebound = p.Rebound
	pi.Assist = p.Assist
	pi.Steal = p.Steal
	pi.Cap = p.Cap
	pi.AppearNum = p.AppearNum
	engine, e := mysql.GetEngine()
	if e != nil {
		logs.Error.Println("db :err ", e)
		return "false", e
	}
	session := engine.NewSession()
	defer session.Close()
	// add Begin() before any action
	session.Begin()
	if _, e1 := session.InsertOne(&pi); e1 != nil {
		logs.Error.Println("db :err1 ", e1)
		return "false", e1
		session.Rollback()
	}
	pe.PlayerId = pi.PlayerId
	pe.DefensiveRebound = p.DefensiveRebound
	pe.Faul = p.Faul
	pe.InsideAttempt = p.InsideAttempt
	pe.InsideMade = p.InsideMade
	pe.OffensiveRebound = p.OffensiveRebound
	pe.ShotAttempt = p.ShotAttempt
	pe.ShotMade = p.ShotMade
	pe.ThreeAttempt = p.ThreeAttempt
	pe.ThreeMade = p.ThreeMade
	pe.TurnOff = p.TurnOff
	pw.BasePower = p.BasePower
	pw.DefensiveRebound = p.DefensiveReboundPower
	pw.InsideAttack = p.InsideAttack
	pw.InsideDefense = p.InsideDefense
	pw.OffensiveRebound = p.OffensiveReboundPower
	pw.OutsideAttack = p.OutsideAttack
	pw.OutsideDefense = p.OutsideDefense
	pw.Pass = p.Pass
	pw.PlayerId = pi.PlayerId
	if _, e2 := session.InsertOne(&pe); e2 != nil {
		logs.Error.Println("db :err2 ", e2)
		return "false", e2
		session.Rollback()
	}
	if _, e3 := session.InsertOne(&pw); e3 != nil {
		logs.Error.Println("db :err2 ", e3)
		return "false", e3
		session.Rollback()
	}
	if err = session.Commit(); err != nil {
		logs.Error.Println("db :err3 ", err)
		return "false", err
	}
	ret = "success"
	return
}

//GetPlayer 获取所有球员
func GetPlayer(limit, start int) (ret []PlayerResult, err error) {

	pis := make([]PlayerInfo, 0)
	pps := make([]PlayerPower, 0)
	pes := make([]PlayerExtro, 0)
	pids := make([]int64, 0)

	pidMap := make(map[int64]PlayerInfo)
	powerMap := make(map[int64]PlayerPower)
	extroMap := make(map[int64]PlayerExtro)
	if engine, e := mysql.GetEngine(); e == nil {
		if e1 := engine.Limit(limit, start).Find(&pis); e1 != nil {
			err = e1
			return
		} else {
			for _, v := range pis {
				pids = append(pids, v.PlayerId)
				pidMap[v.PlayerId] = PlayerInfo{PlayerId: v.PlayerId, Name: v.Name, NickName: v.NickName, Position: v.Position, SecondPosition: v.SecondPosition, Type: v.Type, Score: v.Score, Rebound: v.Rebound, Assist: v.Assist, Steal: v.Steal, Cap: v.Cap, AppearNum: v.AppearNum}
			}
			if e1 := engine.Find(&pps); e1 != nil {
				err = e1
				return
			} else {
				for _, pp := range pps {
					powerMap[pp.PlayerId] = PlayerPower{BasePower: pp.BasePower, OutsideAttack: pp.OutsideAttack, InsideAttack: pp.InsideAttack, OutsideDefense: pp.OutsideDefense, InsideDefense: pp.InsideDefense, OffensiveRebound: pp.OffensiveRebound, DefensiveRebound: pp.DefensiveRebound, Pass: pp.Pass}
				}
			}
			if e1 := engine.Find(&pes); e1 != nil {
				err = e1
				return
			} else {
				for _, pe := range pes {
					extroMap[pe.PlayerId] = PlayerExtro{OffensiveRebound: pe.OffensiveRebound, DefensiveRebound: pe.DefensiveRebound, ThreeAttempt: pe.ThreeAttempt, ThreeMade: pe.ThreeMade, InsideAttempt: pe.InsideAttempt, InsideMade: pe.InsideMade, TurnOff: pe.TurnOff, ShotAttempt: pe.ShotAttempt, ShotMade: pe.ShotMade, Faul: pe.Faul}
				}
			}
		}
	} else {
		logs.Error.Println("db :err ", e)
		return
	}
	for _, v := range pids {
		ret = append(ret, PlayerResult{
			PlayerId:              pidMap[v].PlayerId,
			Name:                  pidMap[v].Name,
			NickName:              pidMap[v].NickName,
			Position:              posEMap[pidMap[v].Position],
			SecondPosition:        posEMap[pidMap[v].SecondPosition],
			Type:                  typeMap[pidMap[v].Type],
			Score:                 pidMap[v].Score,
			Rebound:               pidMap[v].Rebound,
			Assist:                pidMap[v].Assist,
			Steal:                 pidMap[v].Steal,
			Cap:                   pidMap[v].Cap,
			AppearNum:             pidMap[v].AppearNum,
			BasePower:             powerMap[v].BasePower,
			OutsideAttack:         powerMap[v].OutsideAttack,
			InsideAttack:          powerMap[v].InsideAttack,
			OutsideDefense:        powerMap[v].OutsideDefense,
			InsideDefense:         powerMap[v].InsideDefense,
			DefensiveReboundPower: powerMap[v].DefensiveRebound,
			OffensiveReboundPower: powerMap[v].OffensiveRebound,
			Pass:             powerMap[v].Pass,
			OffensiveRebound: extroMap[v].OffensiveRebound,
			DefensiveRebound: extroMap[v].DefensiveRebound,
			ShotAttempt:      extroMap[v].ShotAttempt,
			ShotMade:         extroMap[v].ShotMade,
			ThreeAttempt:     extroMap[v].ThreeAttempt,
			ThreeMade:        extroMap[v].ThreeMade,
			InsideAttempt:    extroMap[v].InsideAttempt,
			InsideMade:       extroMap[v].InsideMade,
			TurnOff:          extroMap[v].TurnOff,
			Faul:             extroMap[v].Faul,
		})
	}
	return
}

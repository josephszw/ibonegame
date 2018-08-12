package models

import (
	"database/sql"
	"fmt"
	"ibgame/logs"
	"ibgame/models/manage_model"
	"ibgame/models/mysql"
	"math"

	"github.com/spf13/cast"

	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

type PlayerInfo struct {
	Player_id       int64 `PK`
	Name            string
	Nick_name       string
	Position        int
	Second_position int
	Type            int
	Score           int
	Rebound         int
	Assist          int
	Steal           int
	Cap             int
	Appear_num      int
}

func GetLink() beedb.Model {
	db, err := sql.Open("mysql", "work:123456@tcp(39.107.94.42:3306)/go?charset=utf8")
	if err != nil {
		panic(err)
	}
	orm := beedb.New(db)
	return orm
}

func GetAll() (playerInfos []PlayerInfo) {
	db := GetLink()
	db.FindAll(&playerInfos)
	return
}
func GetPlayerInfoCount() (ret int64) {
	var pis PlayerInfo
	if engine, e := mysql.GetEngine(); e == nil {
		if count, e1 := engine.Count(&pis); e1 != nil {
			return
		} else {
			ret = count
		}
	} else {
		logs.Error.Println("db :err ", e)
		return
	}
	return
}
func GetPlayerInfoLimit(start, end int) (playerInfo []manage_model.PlayerResult) {
	playerInfo, _ = manage_model.GetPlayer(end, start)
	return
}

func GetPlayerInfo(id int64) (playerInfo PlayerInfo) {
	db := GetLink()
	db.Where("player_id=?", id).Find(&playerInfo)
	return
}

func SavePlayerInfo(playerInfo PlayerInfo) (bg PlayerInfo) {
	fmt.Println(playerInfo)
	db := GetLink()
	playerInfo.Player_id = cast.ToInt64(playerInfo.Player_id)
	err := db.Save(&playerInfo)
	fmt.Println(err)
	return bg
}

func DelPlayerInfo(playerInfo PlayerInfo) {
	db := GetLink()
	db.Delete(&playerInfo)
	return
}

//分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
func Paginator(page, prepage int, nums int64) map[string]interface{} {

	var firstpage int //前一页地址
	var lastpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		firstpage = page - 1
		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstpage = int(math.Max(float64(1), float64(page-1)))
		lastpage = page + 1
		//fmt.Println(pages)
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["lastpage"] = lastpage
	paginatorMap["currpage"] = page
	return paginatorMap
}

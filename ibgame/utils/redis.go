package utils

// import (
// 	"ibgame/logs"

// 	"github.com/garyburd/redigo/redis"
// )

// var conn redis.Conn

// func init() {
// 	var err error
// 	conn, err = redis.Dial("tcp", "39.107.94.42:6379")
// 	if err != nil {
// 		logs.Error.Println("Connect to redis error", err)
// 		return
// 	}
// 	defer conn.Close()
// }

// func Set(filed string, value int) (err error) {
// 	_, err = conn.Do("set", filed, value)
// 	if err != nil {
// 		logs.Error.Println("redis set err")
// 	}
// 	return
// }

// func Get(filed string) (ret int, err error) {
// 	ret, err = redis.Int(conn.Do("get", filed))
// 	if err != nil {
// 		logs.Error.Println("redis get err")
// 	}
// 	return
// }

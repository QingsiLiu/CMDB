package conn

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func GetConn() redis.Conn {
	conn, err := redis.Dial("tcp", "110.42.144.237:6379", redis.DialPassword("root"))
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

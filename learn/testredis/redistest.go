package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

func main() {

	conn, err := redis.Dial("tcp", "110.42.144.237:6379", redis.DialPassword("root"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	keys, _ := redis.Strings(conn.Do("KEYS", "*"))
	fmt.Println(keys)

	conn.Do("SET", "cmdb:test:starttime", time.Now().Format("2006-01-02 15:04:05"))

	value, _ := redis.String(conn.Do("GET", "cmdb:test:starttime"))
	fmt.Println(value)

	for i := 0; i < 10; i++ {
		conn.Do("LPUSH", "cmdb:test:tasks", i)
	}

	/*for i:=0; i<5; i++ {
		fmt.Println(redis.Strings(conn.Do("BRPOP", "cmdb:test:tasks", 0)))
	}*/

	conn.Do("HMSET", redis.Args{}.Add("cmdb:test:user").Add("name").Add("lhq")...)
	user, _ := redis.StringMap(conn.Do("HGETALL", "cmdb:test:user"))
	fmt.Println(user)

}

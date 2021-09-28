package consumer

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"log"
	"magego/course-33/cmdb/learn/testredis/testqueue/conn"
	"time"
)

func Run() {
	rConn := conn.GetConn()
	defer rConn.Close()

	lockKey := "consumer:master:uid"
	uid := uuid.New().String()
	log.Println("uid", uid)

	go func() {
		rConn := conn.GetConn()
		defer rConn.Close()
		ticker := time.NewTicker(time.Second * 10)
		defer ticker.Stop()
		for range time.Tick(time.Second * 10) {
			rt, err := rConn.Do("SET", lockKey, uid, "EX", 30, "NX")
			if err != nil || rt != "OK" {
				rt, err := redis.String(rConn.Do("GET", lockKey))
				if err == nil {
					if rt == uid {
						log.Printf("is master 2")
						// 续时间
						rConn.Do("Expire", lockKey, 30)
					} else {
						log.Printf("not master")
					}
				}
			} else {
				log.Printf("is master 1")
			}
			<-ticker.C
		}
	}()

	for {
		if rt, err := redis.String(rConn.Do("Get", lockKey)); err != nil || rt == uid {
			log.Printf("is not master, no worker")
			time.Sleep(time.Second * 10)
			continue
		}

		values, err := redis.Strings(rConn.Do("BRPOP", "cmdb:test:works", 3))
		if err != nil {
			continue
		}
		fmt.Println(values[1])
	}

}

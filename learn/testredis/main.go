package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"math/rand"
	"os"
	"time"
)

func getConn() redis.Conn {
	conn, err := redis.Dial("tcp", "110.42.144.237:6379", redis.DialPassword("root"))
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func producer() {
	// 随机sleep n s 在队列cmdb:test:works 放入当时时间

	conn := getConn()
	defer conn.Close()

	for {
		t := time.Now().Format("2006-01-02 15:04:05")

		conn.Do("LPUSH", "cmdb:test:works", fmt.Sprintf("%d:%s", os.Getpid(), t))
		time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	}
}

func consumer() {
	conn := getConn()
	defer conn.Close()

	for {
		values, err := redis.Strings(conn.Do("BRPOP", "cmdb:test:works", 3))
		if err != nil {
			continue
		}
		fmt.Println(values[1])
	}

}

func main() {

	rand.Seed(time.Now().Unix())
	go consumer() // 消费者
	producer()    // 生产者

	// 消费者只能有一个
	// 分布式锁
	// set lock value ex 30 nx // value随机（一个进程随机）
	// go 10s set
	// get lock == uuid 执行
	// 设置失败或者get lock != uuid 就不执行
}

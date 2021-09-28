package main

import (
	"flag"
	"fmt"
	"magego/course-33/cmdb/learn/testredis/testqueue/consumer"
	"magego/course-33/cmdb/learn/testredis/testqueue/producer"
)

func main() {
	var isProducer bool
	flag.BoolVar(&isProducer, "p", false, "producer")
	flag.Parse()
	if isProducer {
		//生产者
		fmt.Println("生产者")
		producer.Run()
	} else {
		// 消费者
		fmt.Println("消费者")
		consumer.Run()
	}
}

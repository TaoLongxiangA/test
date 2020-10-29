package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	result, err := client.Ping().Result()

	if err != nil {
		log.Println("ping err:", err)
	}
	log.Println("ping success:", result)

	for i := 0; i < 5; i++ {
		err = client.LPush("age", i).Err()
		if err != nil {
			log.Printf("lpush req failed. Error : %v", err)
		}
	}

	log.Printf("test")

	for {
		strings, err := client.BRPop(time.Minute, "age").Result()
		if err != nil {
			continue
		}
		fmt.Println(strings[0])
		fmt.Println(strings[1])
	}
}

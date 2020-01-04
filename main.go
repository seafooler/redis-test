package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:"127.0.0.1:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	//for i:=0; i < 10000; i++ {
		/*for j:=0; j < i+1; j++ {
			//client.SetBit(strconv.Itoa(i), 1, 1)
			//client.SetBit(strconv.Itoa(i), 2, 1)
			//client.SetBit(strconv.Itoa(i), 3, 1)
			//client.SetBit(strconv.Itoa(i), 4, 1)
			//client.SetBit(strconv.Itoa(i), 5, 1)
			//client.SetBit(strconv.Itoa(i), 6, 1)
			//client.SetBit(strconv.Itoa(i), 7, 1)
			//client.SetBit(strconv.Itoa(i), 8, 1)
			//client.SetBit(strconv.Itoa(i), 9, 1)
			//client.SetBit(strconv.Itoa(i), 10, 1)
			//client.SetBit(strconv.Itoa(i), 11, 1)
			//client.SetBit(strconv.Itoa(i), 21, 1)
			//client.SetBit(strconv.Itoa(i), 12, 1)
			//client.SetBit(strconv.Itoa(i), 13, 1)
			//client.SetBit(strconv.Itoa(i), 14, 1)
			//client.SetBit(strconv.Itoa(i), 15, 1)
			//client.SetBit(strconv.Itoa(i), 16, 1)
			//client.SetBit(strconv.Itoa(i), 17, 1)
			//client.SetBit(strconv.Itoa(i), 18, 1)
			//client.SetBit(strconv.Itoa(i), 19, 1)
			//client.SetBit(strconv.Itoa(i), 100, 1)
			//client.SetBit(strconv.Itoa(i), 1000, 1)
			//client.SetBit(strconv.Itoa(i), 8000, 1)
			fmt.Printf("%d, %d\n", i, j)
			client.SetBit(strconv.Itoa(i), int64(j), 1)
		}*/
		//client.SetBit(strconv.Itoa(i), int64(1), 1)
	//}

	//client.SetBit(strconv.Itoa(1), int64(100000000), 1)

	start := time.Now()
	for i:=0; i< 10000; i++ {
		client.GetBit("1", 1)
		/*_, err := client.GetBit("1", int64(i)).Result()
		if err != nil {
			log.Fatal(err)
		}*/
		//fmt.Println(value)
	}

	fmt.Println("Fetch bit in redis: ", time.Now().Sub(start).String())

}

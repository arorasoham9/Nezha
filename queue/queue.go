package queue

import(
	"github.com/go-redis/redis"
	"fmt"


)



func StartQueue(PORT string) {
	//begin client
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:"+PORT,
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	//check if a container exists

}
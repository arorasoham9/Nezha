package main

import (
	// ssh "/ssh"
	// b "ECE49595_PROJECT/basic"
	// ssh "ECE49595_PROJECT/ssh"
	q "ECE49595_PROJECT/queue"

	"github.com/go-redis/redis"
	"fmt"
)



func main(){

	queue := q.MakeQueue(&redis.Options{ 
		Addr: "localhost:6379", 
		Password: "", 
		 DB: 0, 
   }, &redis.Options{ 
	Addr: "localhost:6379", 
	Password: "", 
	 DB: 0, 
} )
	q.AddRequestToQueue(queue, "test9", q.Queue_Request{NAME: "sabash", 
	EMAIL: "sabutdana@gmail.com",
	CURRENT_IP: "mei nahi bataonga",
	LOCATION: "jhadio k peeche",
	CREATED_AT: "cake murder day", 
	LASTSEEN:"don ko dhundna mushkil hi nahi namumkin hai"})
	fmt.Println(q.GetRequestFromQueue(queue, "test9"))

}
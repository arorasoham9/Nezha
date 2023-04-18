package main

import (
	// ssh "/ssh"
	// b "ECE49595_PROJECT/basic"
	// ssh "ECE49595_PROJECT/ssh"
	q "ECE49595_PROJECT/queue"
	// // "context"
	d "ECE49595_PROJECT/dock"
	"github.com/go-redis/redis"
	// "fmt"
	// "time"
)

func main(){
	d.InitDock()
	d.StopAllContainers()
// 	fmt.Println( q.MakeQueue(&redis.Options{ 
// 		Addr: "localhost:6379", 
// 		Password: "", 
// 		 DB: 0, 
//    }, &redis.Options{ 
// 	Addr: "localhost:6379", 
// 	Password: "", 
// 	 DB: 0, 
// } ))

	// fmt.Println(q.Queue_container_name)
	// err := q.AddRequestToQueue( "test94", q.Queue_Request{NAME: "sabash", 
	// EMAIL: "sabutdana@gmail.com",
	// CURRENT_IP: "mei nahi bataonga",
	// LOCATION: "jhadio k peeche",
	// CREATED_AT: "cake murder day", 
	// LASTSEEN:"don ko dhundna mushkil hi nahi namumkin hai"})
	
	// if err != nil{
	// 	fmt.Println(err)
	// }

	// fmt.Println(q.GetRequestFromQueue( "test94"))
	// // fmt.Println(q.RestartQueue(false))
	// fmt.Println(q.ShutDownQueue(false))
	// // fmt.Println(q.GetRequestFromQueue( "test94"))

	q.BeginQueueOperation(&redis.Options{ 
		Addr: "localhost:6379", 
		Password: "", 
		 DB: 0, 
   }, &redis.Options{ 
	Addr: "localhost:6379", 
	Password: "", 
	 DB: 0, 
}, 5, 5000)


}


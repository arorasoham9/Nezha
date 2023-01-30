package main

import (
	// ssh "/ssh"
	// b "ECE49595_PROJECT/basic"
	// ssh "ECE49595_PROJECT/ssh"
	q "ECE49595_PROJECT/queue"
	"fmt"
)
const address = "localhost"
const port = "6379"
var user string   = "root"
var host_port  string =  "45.79.26.72:22"
var cmd  string = "uname -a"
var  pass  string = "BYe7_c9p.RQYhew"



func main(){

	conn, err := q.MakeQueueConnection(port, address, "", 0, q.API_Q_CLI )
	if err != nil{
		fmt.Println("successful")
	}
	fmt.Println(conn.Ping().Result())
}
package main

import (
	// ssh "/ssh"
	// b "ECE49595_PROJECT/basic"
	ssh "ECE49595_PROJECT/ssh"
	q "ECE49595_PROJECT/queue"
)

var user string   = "root"
var host_port  string =  "45.79.26.72:22"
var cmd  string = "uname -a"
var  pass  string = "BYe7_c9p.RQYhew"



func main(){

	ssh.RunCli(user, host_port, cmd, pass)
	q.StartQueue("9999")
}
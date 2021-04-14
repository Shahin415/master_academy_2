package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	//net = network
	nl, err := net.Listen("tcp", ":8888") //ip:port  192.168.0.102
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //1

	}

	conn, err := nl.Accept() //Layer-5 session layer
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //1
	}

	remoteAddr := conn.RemoteAddr().String()
	fmt.Println(remoteAddr)

	//byte
	conn.Write([]byte("welcome we have received your message"))

	conn.Close() //red button press
	nl.Close()   //red button press

}

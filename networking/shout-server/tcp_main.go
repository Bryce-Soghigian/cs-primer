package main

import (
	"fmt"
	"net"
)


 func main() {
	fmt.Println("Starting the tcp shout server ...")	
	ln, err := net.Listen("tcp", ":8080") 
	if err != nil { 
		panic(err) 
	} 

	for { 
		conn, err := ln.Accept() 
		if err != nil { 
			panic(err) 
		} 

		go handleConnection(conn) 
	}

}


func handleConnection(conn net.Conn) { 
	defer conn.Close() 

	for { 
		buf := make([]byte, 1024) 
		n, err := conn.Read(buf) 
		if err != nil { 
			panic(err) 
		} 

		println(string(buf[:n])) 

		_, err = conn.Write([]byte("Message received\n")) 
		if err != nil { 
			panic(err) 
		} 
	} 
}

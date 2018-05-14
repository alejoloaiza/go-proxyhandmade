/*
   Socks5 proxy server by golang
   http://github.com/ring04h/s5.go
   reference: shadowsocks go local.go
   https://github.com/shadowsocks/shadowsocks-go
*/

package main

import (
	"flag"
	"fmt"
	"go-proxyhandmade/connection"
	"log"
	"net"
	"runtime"
)

func main() {

	// maxout concurrency
	runtime.GOMAXPROCS(runtime.NumCPU())

	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8080", "proxy listen address")
	flag.Parse()

	connection.Verbose = *verbose

	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(err)
		return
	}

	log.Printf("Listening %s \n", *addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error in Accept: " + err.Error())
			return
		}
		if connection.Verbose {
			log.Println("new client:", conn.RemoteAddr())
		}
		go connection.HandleConnection(conn)
	}
}

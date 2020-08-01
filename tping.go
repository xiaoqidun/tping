package main

import (
	"flag"
	"log"
	"net"
	"time"
)

func main() {
	h := flag.String("h", "aite.xyz", "")
	p := flag.String("p", "443", "")
	flag.Usage = func() {}
	flag.Parse()
	address := net.JoinHostPort(*h, *p)
	if n, err := TcpPing(address); err == nil {
		log.Println("地址", address, "延迟", n/1e6, "毫秒")
	} else {
		log.Println("地址", address, "错误", err)
	}
}

func TcpPing(address string) (n int64, err error) {
	s := time.Now()
	tcpConn, err := net.Dial("tcp", address)
	n = time.Now().Sub(s).Nanoseconds()
	if err != nil {
		return
	}
	tcpConn.Close()
	return
}

package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

var w *sync.WaitGroup

func main() {
	w = &sync.WaitGroup{}
	for i := 2; i < 255; i++ {
		addr := "192.168.1." + strconv.Itoa(i) + ":22"
		w.Add(1)
		go connect(addr)
	}
	w.Wait()
}

func connect(addr string) {
	defer w.Done()
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		//error nothing to do
	} else {
		fmt.Println("OPEN", addr)
		conn.Close()
	}
}

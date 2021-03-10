package main

import (
	"flag"
	"strconv"
)

func main() {
	addr := flag.String("addr", "localhost", "bind ip")
	port := flag.Int("port", 5555, "bind port")
	delim := flag.String("delim", ":", "What delimeter do you want to use?")
	flag.Parse()

	servStr := *addr + ":" + strconv.Itoa(*port)
	tcpServ, err := NewServer(servStr, *delim)

	if err != nil {
		panic(err)
	}

	tcpServ.Run()
}

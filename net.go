package main

import (
	"bufio"
	"errors"
	"io"
	"net"
	"net/textproto"
	"os"

	log "github.com/sirupsen/logrus"
)

type TcpServer struct {
	addr      string
	server    net.Listener
	processor Processor
}

type Server interface {
	Run() error
	Close() error
}

func NewServer(addr string, delim string) (Server, error) {
	p := Processor{
		writer: os.Stdout,
		delim:  delim,
	}
	return &TcpServer{
		addr:      addr,
		processor: p,
	}, nil
}

func (t *TcpServer) Run() (err error) {
	t.server, err = net.Listen("tcp", t.addr)
	if err != nil {
		return err
	}
	defer t.Close()
	log.Info("Listening for connections on: ", t.addr+"...")

	for {
		conn, err := t.server.Accept()
		if err != nil {
			err = errors.New("could not accept connection")
			break
		}
		if conn == nil {
			err = errors.New("Could not create connection")
			break
		}
		log.Info("New connection from: ", conn.RemoteAddr().String())
		go t.handleConnection(conn)
	}
	return
}

func (t *TcpServer) Close() (err error) {
	return t.server.Close()
}

func (t *TcpServer) handleConnection(conn net.Conn) error {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	// read data line by line from the socket
	for {
		line, err := tp.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			break
		}
		t.processor.processLine(line)

	}
	return nil
}

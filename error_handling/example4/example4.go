package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

//example4 provides code to show how to implement behavior as context

type client struct {
	name   string
	reader *bufio.Reader
}

func (c *client) TypeAsContext() {
	for {
		line, err := c.reader.ReadString('\n')
		if err != nil {
			switch e := err.(type) {
			case *net.OpError:
				if !e.Temporary() {
					log.Println("Temporary:Client leaving chat")
					return
				}
			case *net.AddrError:
				if !e.Temporary() {
					log.Println("Temporary:Client leaving chat")
					return
				}
			case *net.DNSConfigError:
				if !e.Temporary() {
					log.Println("Temporary:Client leaving chat")
					return
				}
			default:
				if err == io.EOF {
					log.Println("EOF:Client leaving chat")
					return
				}
			}
		}
		fmt.Println(line)
	}
}

type temporary interface {
	Temporary() bool
}

func (c *client) BehaviorAsContext() {
	for {
		line, err := c.reader.ReadString('\n')
		if err != nil {
			switch e := err.(type) {
			case temporary:
				if !e.Temporary() {
					log.Println("Tempoary: Client leaving chat")
					return
				}
			default:
				if err == io.EOF {
					log.Println("EOF: Client leaving chat")
					return
				}
				log.Println("reading-routline", err)

			}
		}
		fmt.Println(line)
	}
}

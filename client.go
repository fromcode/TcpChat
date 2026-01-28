package main

import "net"

type client struct {
	conn     net.Conn
	nick     string
	room     *room // mengambil dari room command
	commands chan<- command
}

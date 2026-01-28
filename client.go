package main

import "net"

type client struct {
	conn    net.Addr
	nick    string
	room    *room // mengambil dari room command
	command chan<- commandID
}

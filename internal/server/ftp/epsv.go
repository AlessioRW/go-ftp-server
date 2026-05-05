package ftp

import (
	"fmt"
	"log"
	"net"
)

func (c *Conn) epsv() {
	if c.passiveListener != nil {
		c.closePassiveListener()
	}

	srv, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Print("ERROR failed to create passive listener:  ", err)
		c.respond(status425)
		return
	}
	c.passiveListener = srv

	newPort := srv.Addr().(*net.TCPAddr).Port
	res := fmt.Sprintf(status229, newPort)
	c.respond(res)
}

package ftp

import (
	"fmt"
	"log"
	"net"
)

func (d *dataPort) toAddress() string {
	if d == nil {
		return ""
	}

	return fmt.Sprintf("%v:%v", d.ip, d.port)

}

func (c *Conn) dataConnect() (net.Conn, error) {
	if c.passiveListener != nil { // use passive listener if exists
		conn, err := c.passiveListener.Accept()
		if err != nil {
			log.Print("ERROR failed to get passive listener connection: ", err)
			return nil, err
		}
		c.closePassiveListener()
		return conn, nil
	}

	conn, err := net.Dial("tcp", c.dataPort.toAddress())
	if err != nil {
		log.Print("ERROR failed to dial client: ", err)
		return nil, err
	}
	return conn, nil
}

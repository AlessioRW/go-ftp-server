package ftp

import (
	"log"
	"net"
)

// Conn represents a connection to the FTP server
type Conn struct {
	conn            net.Conn
	dataPort        *dataPort
	passiveListener net.Listener // second server/listener for passive mode

	dataType dataType

	rootDir string // root dir of music server
	workDir string // client conn working dir, allows per client navigation
}

func NewConn(c net.Conn, path string) *Conn {
	return &Conn{
		conn:    c,
		rootDir: path,
		workDir: "/",
	}
}

func (c *Conn) closePassiveListener() {
	err := c.passiveListener.Close()
	if err != nil {
		log.Print("error closing passive listener: ", err)
	}
	c.passiveListener = nil
}

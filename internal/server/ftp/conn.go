package ftp

import (
	"net"
)

// Conn represents a connection to the FTP server
type Conn struct {
	conn     net.Conn
	dataType dataType
	dataPort *dataPort
	rootDir  string // root dir of music server
	workDir  string // client conn working dir, allows per client navigation
}

func NewConn(c net.Conn, path string) *Conn {
	return &Conn{
		conn:    c,
		rootDir: path,
		workDir: "/",
	}
}

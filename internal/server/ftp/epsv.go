package ftp

import (
	"fmt"
	"log/slog"
	"net"
)

// TODO: read how EPSV actually provides the connection
// currently assuming it just opens a new ftp server
func (c *Conn) epsv() {
	if c.passiveListener != nil {
		c.passiveListener.Close()
		c.passiveListener = nil
	}

	srv, err := net.Listen("tcp", ":0")
	if err != nil {
		slog.Error("failed to create passive listener", "error", err)
		c.respond(status425)
		return
	}
	c.passiveListener = srv

	newPort := srv.Addr().(*net.TCPAddr).Port
	res := fmt.Sprintf(status229, newPort)
	c.respond(res)
}

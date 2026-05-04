package ftp

import (
	"fmt"
	"log/slog"
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
		c.passiveListener.Close()
		c.passiveListener = nil
		if err != nil {
			slog.Error("failed to get passive listener connection", "error", err)
			return nil, err
		}
		return conn, nil
	}

	conn, err := net.Dial("tcp", c.dataPort.toAddress())
	if err != nil {
		slog.Error("failed to dial client", "error", err)
		return nil, err
	}
	return conn, nil
}

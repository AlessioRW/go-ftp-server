package ftp

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func (c *Conn) pasv() {
	if c.passiveListener != nil {
		c.closePassiveListener()
	}

	srv, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Print("ERROR failed to create passive listener: ", err)
		c.respond(status425)
		return
	}
	c.passiveListener = srv

	// TODO: add code to determine IP
	ip := "127.0.0.1"
	sIP := strings.Split(ip, ".")

	newPort := srv.Addr().(*net.TCPAddr).Port
	p1, p2, err := getPortSegments(newPort)
	if err != nil {
		c.respond(status451)
		return
	}
	res := fmt.Sprintf(
		status227,
		sIP[0], sIP[1], sIP[2], sIP[3],
		p1, p2,
	)
	c.respond(res)
}

func getPortSegments(port int) (int, int, error) {
	hex := fmt.Sprintf("%x", port)
	p1, err := strconv.ParseInt(hex[0:2], 16, 64)
	if err != nil {
		log.Print("ERROR failed to parse port into int: ", err)
		return -1, -1, err
	}
	p2, err := strconv.ParseInt(hex[2:], 16, 64)
	if err != nil {
		log.Print("ERROR failed to parse port into int: ", err)
		return -1, -1, err
	}

	return int(p1), int(p2), nil
}

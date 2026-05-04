package ftp

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

type dataPort struct {
	ip   string
	port int
}

// port args is []string, length=6
//   - args[0:3] = ip addr as bytes
//   - args[4:] = port as bytes, calc as `(256*args[4]) + args[5]`
func (c *Conn) port(args []string) {
	if len(args) != 1 {
		c.respond(status501)
		return
	}
	// slog.Info("DEBUG::PORT", "args", args)

	dp, err := dataPortFromArgs(strings.Split(args[0], ","))
	if err != nil {
		c.respond(status501)
		return
	}

	c.dataPort = dp
	c.respond(status200)

}

func dataPortFromArgs(addr []string) (*dataPort, error) {
	var dp dataPort

	if len(addr) != 6 {
		err := errors.New("invlaid address argument, length should be 6")
		slog.Error("failed to form data port", "error", err)
		return nil, err
	}

	ip := fmt.Sprintf(
		"%v.%v.%v.%v",
		addr[0], addr[1], addr[2], addr[3],
	)
	dp.ip = ip

	p1, err := strconv.Atoi(addr[4])
	if err != nil {
		slog.Error("failed to convert port byte 1 to int", "error", err)
		return nil, err
	}
	p2, err := strconv.Atoi(addr[5])
	if err != nil {
		slog.Error("failed to convert port byte 2 to int", "error", err)
		return nil, err
	}
	dp.port = (p1 * 256) + p2

	return &dp, nil
}

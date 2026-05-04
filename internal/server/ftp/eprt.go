package ftp

import (
	"log/slog"
	"strconv"
	"strings"
)

// eprt args are in format of:
// [ IP_TYPE(v4/v6) | IP | PORT ]
// not sure if divider is always "|" (is for curl)
func (c *Conn) eprt(args []string) {
	splitArgs := filterEmpty(strings.Split(args[0], "|"))
	if len(splitArgs) != 3 {
		slog.Error("eprt command does not have correct number of args")
		c.respond(status501)
		return
	}

	var dp *dataPort
	var err error

	switch splitArgs[0] {
	case "1":
		dp, err = parseIPv4(splitArgs[1:])
		if err != nil {
			c.respond(status501)
			return
		}
	case "2":
		// TODO: implement IPv6
		// dp, err = parseIPv6(splitArgs[1:])
		// if err != nil {
		// 	c.respond(status501)
		// 	return
		// }
		c.respond(status504)
	default:
		c.respond(status501)
	}

	c.dataPort = dp
	c.respond(status200)
}

func parseIPv4(args []string) (*dataPort, error) {
	var dp dataPort

	// TODO: validate IP format properly?
	dp.ip = args[0]

	port, err := strconv.Atoi(args[1])
	if err != nil {
		slog.Error("failed to convert port to int", "error", err)
		return nil, err
	}
	dp.port = port

	return &dp, nil
}

// remove empty strings from list
func filterEmpty(args []string) []string {
	nArgs := []string{}
	for _, s := range args {
		if strings.TrimSpace(s) != "" {
			nArgs = append(nArgs, s)
		}
	}
	return nArgs
}

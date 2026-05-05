package ftp

import (
	"log"
	"strconv"
	"strings"
)

// eprt args are in format of:
// [ IP_TYPE(v4/v6) | IP | PORT ]
// not sure if divider is always "|" (is for curl)
func (c *Conn) eprt(args []string) {
	splitArgs := filterEmpty(strings.Split(args[0], "|"))
	if len(splitArgs) != 3 {
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
		dp, err = parseIPv6(splitArgs[1:])
		if err != nil {
			c.respond(status501)
			return
		}
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
		log.Print("ERROR failed to convert port to int: ", err)
		return nil, err
	}
	dp.port = port

	return &dp, nil
}

func parseIPv6(args []string) (*dataPort, error) {
	var dp dataPort

	// TODO: handle more than localhost
	dp.ip = "::1"

	port, err := strconv.Atoi(args[1])
	if err != nil {
		log.Print("ERROR failed to convert port to int", err)
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

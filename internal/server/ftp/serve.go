package ftp

import (
	"bufio"
	"log/slog"
	"strings"
)

func Serve(c *Conn) {
	// send OK to client
	c.respond(status220)

	// loop over tcp conn and listen for commands
	s := bufio.NewScanner(c.conn)
	for s.Scan() {
		input := strings.Fields(s.Text())
		if len(input) == 0 {
			continue
		}

		command, args := input[0], input[1:]
		slog.Info(
			"recieved command",
			"command", command,
			"args", args,
		)

		switch command {
		case "CWD": // cd with client connection
			c.cwd(args)
		case "PWD": // return working dir
			c.pwd()
		case "LIST": // ls files in wd, may need different implementation
			c.list(args)
		case "NLST": // simple ls files in wd
			c.list(args)
		case "PORT": // sets client's data channel address
			c.port(args)
		case "EPRT": // sets client's data channel address, supports ipv6
			c.eprt(args)
		case "USER": // this just prints the passed user, not too useful?
			c.user(args)
		case "RETR": // send client file
			c.retr(args)
		case "SIZE": // returns file size
			c.size(args)
		case "TYPE": // set data type expected by client
			c.setDataType(args)
		case "PASV": // set client to use passive mode (NOTE: find client to test PASV with)
			c.respond(status502)
			// c.pasv(args)
		case "EPSV": // set client to use extended passive mode
			c.epsv()
		case "QUIT": // close connection
			c.respond(status221)
			return

		default: // command not implented
			c.respond(status502)
		}
	}

	if s.Err() != nil {
		slog.Error(
			"error scanning connection",
			"errror", s.Err(),
		)
	}
}

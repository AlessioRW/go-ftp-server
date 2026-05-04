package ftp

import (
	"fmt"
	"log/slog"
)

// standard FTP responce codes
// see https://en.wikipedia.org/wiki/List_of_FTP_server_return_codes
const (
	status150 = "150 command status okay; about to open data connection."
	status200 = "200 Command okay."
	status213 = "213 %v"
	status220 = "220 Service ready for new user."
	status221 = "221 Service closing control connection."
	status226 = "226 Closing data connection. Requested file action successful."
	status229 = "229 Entering Extended Passive Mode (|||%v|)."
	status230 = "230 User %s logged in, proceed."
	status257 = "257 \"%v\" is the current directory"
	status425 = "425 Can't open data connection."
	status426 = "426 Connection closed; transfer aborted."
	status451 = "451 Requested action aborted. Local error in processing."
	status501 = "501 Syntax error in parameters or arguments."
	status502 = "502 Command not implemented."
	status504 = "504 Cammand not implemented for that parameter."
	status550 = "550 Requested action not taken. File unavailable."
)

func (c *Conn) respond(s string) {
	slog.Info(
		"sending response",
		"response", s,
	)

	// write response to client (via net connection writer)
	// net package nicely handles this for you :)
	_, err := fmt.Fprint(c.conn, s, c.EOL())

	if err != nil {
		slog.Error(
			"failed to send response",
			"response", s,
			"error", err,
		)
	}
}

func (c *Conn) EOL() string {
	switch c.dataType {
	case ascii:
		return "\r\n"
	case binary:
		return "\n"
	default:
		return "\n"
	}
}

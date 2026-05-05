package ftp

import (
	"fmt"
	"log"
)

// standard FTP responce codes
// see https://en.wikipedia.org/wiki/List_of_FTP_server_return_codes
const (
	status125 = "125 Data connection already open"
	status150 = "150 Opening data connection"
	status200 = "200 Command okay."
	status211 = "211-%v211 End" // feat, stat (specified following hypen)
	status213 = "213 %v"        // size
	status215 = "215 %v"        // syst
	status220 = "220 Service ready for new user."
	status221 = "221 Service closing control connection."
	status226 = "226 Requested file action successful"
	status227 = "227 Entering Passive Mode (%v,%v,%v,%v,%v,%v)."
	status229 = "229 Entering Extended Passive Mode (|||%v|)."
	status230 = "230 User %s logged in, proceed."
	status250 = "250 Directory successfully changed."
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
	log.Print("response >> ", s)

	// write response to client (via net connection writer)
	// net package nicely handles this for you :)
	_, err := fmt.Fprint(c.conn, s, c.EOL())
	if err != nil {
		log.Print("ERROR failed to send response: ", err)
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

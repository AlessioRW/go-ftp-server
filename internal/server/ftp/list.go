package ftp

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

const LIST_TIME_FORMAT = "Jan _2 15:04"

func (c *Conn) list(args []string) {
	var target string
	var dir string
	flags := []string{}

	// loop over args and remove flags
	for _, arg := range args {
		if arg[0] == '-' {
			flags = append(flags, arg)
		} else {
			dir = arg
		}
	}

	// form list target
	if dir != "" {
		target = filepath.Join(c.rootDir, c.workDir, dir)
	} else {
		target = filepath.Join(c.rootDir, c.workDir)
	}

	files, err := os.ReadDir(target)
	if err != nil {
		log.Print("ERROR failed to read directory '"+target+"': ", err)
		c.respond(status550)
		return
	}

	// open connection to client
	// used to write files back as a response
	dataConn, err := c.dataConnect()
	if err != nil {
		c.respond(status425)
		return
	}
	defer dataConn.Close()

	// send 150 ack once data connection established
	c.respond(status150)

	// loop over files and write to client
	for _, file := range files {
		fileData, err := file.Info()
		if err != nil {
			log.Print("ERROR failed to get file data", "file", file.Name(), "error", err)
			c.respond(status451)
			return
		}

		_, err = fmt.Fprint(dataConn, formatFileData(fileData), c.EOL())
		if err != nil {
			log.Print("ERROR failed to write file name: ", err)
			c.respond(status451)
			return
			// should command be terminated here??
			// code 426 specifed transfer aborted but example code doesn't
		}
	}

	_, err = fmt.Fprint(dataConn, c.EOL())
	if err != nil {
		log.Print("ERROR failed to write closing line: ", err)
		c.respond(status426)
		return
	}

	c.respond(status226)
}

func formatFileData(f fs.FileInfo) string {
	return fmt.Sprintf(
		"%s %d %s %s %d %s %s %s %s",
		f.Mode().String(),           // permissions
		1,                           // link count
		"_",                         // owner
		"_",                         // group
		f.Size(),                    // size
		f.ModTime().Format("Jan"),   // mmm
		f.ModTime().Format("02"),    // dd
		f.ModTime().Format("15:04"), // hh:mm
		f.Name(),                    // filename
	)
}

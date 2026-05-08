package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type config struct {
	StorageRoot string
	HostPort    int
}

var Config config

func Parse(args []string) error {
	conf := config{
		HostPort:    21,
		StorageRoot: "",
	}
	for i, arg := range args {
		if i == len(args)-1 {
			break
		}

		switch arg {
		case "--storage", "-s":
			d, err := os.Stat(args[i+1])
			if err != nil {
				log.Printf("ERROR argument %v, unable to verify \"%v\" exists", arg, args[i+1])
				return fmt.Errorf("ERROR argument %v", arg)
			}
			if !d.IsDir() {
				log.Printf("ERROR argument %v, \"%v\" is not a folder", arg, args[i+1])
				return fmt.Errorf("ERROR argument %v", arg)
			}
			conf.StorageRoot = args[i+1]
		case "--port", "-p":
			port, err := strconv.Atoi(args[i+1])
			if err != nil {
				log.Printf("ERROR argument %v, unable to parse \"%v\" as a number", arg, args[i+1])
				return fmt.Errorf("ERROR argument %v", arg)
			}
			conf.HostPort = port
		}
	}
	Config = conf

	return nil
}

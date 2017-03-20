package syncman

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	args := make([]string, 0)
	if len(os.Args) > 0 {
		args = os.Args[1:]
	}
	s := parseFlags(args)

	sm := newSyncManager(s)
	sm.Run()
}

func parseFlags(args []string) *settings {
	s := settings{}
	fs := flag.NewFlagSet("syncman", flag.ExitOnError)
	fs.UintVar(&s.Port, "port", 6000, "port number")
	fs.StringVar(&s.ConfigPath, "config", "config.xml", "path to config file")
	fs.Parse(args)
	return &s
}

type settings struct {
	Port       uint
	ConfigPath string
}

type syncManager struct {
	settings *settings
}

func newSyncManager(s *settings) *syncManager {
	sm := syncManager{settings: s}
	return &sm
}

func (sm *syncManager) Run() error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", sm.settings.Port))
	if err != nil {
		return err
	}

	for {
		rw, err := ln.Accept()
		// TODO design and test error handling
		if err != nil {
			log.Println(err)
		}

		// handle
		scanner := bufio.NewScanner(rw)
		ok := scanner.Scan()
		if !ok {
			continue
		}
		line := scanner.Text()
		var resp string
		if line == "+PING\r\n" {
			resp = "+OK\r\n"
		} else {
			resp = "-Unknown command: "
		}
		_ = resp
	}
}

package syncman

import (
	"flag"
	"os"
)

func main() {
	args := make([]string, 0)
	if len(os.Args) > 0 {
		args = os.Args[1:]
	}
	_ = parseFlags(args)
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

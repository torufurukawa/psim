package syncman

import "flag"

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

package main

import "get-block/config"

var cnf = config.Config{}

func init() {
	cnf.ReadConfig("./var/config")
}

func main() {
	application := new(Application)
	application.Run(&cnf)
}

package main

import (
	"flag"
	"golang-gin-boilerplate/server"
)

func main() {
	mode := flag.String("mode", "development", "For environtment variables")
	flag.Parse()

	server.Init(*mode)
}

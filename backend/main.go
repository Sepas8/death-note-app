package main

import (
	"backend-avanzada/server"
)

func main() {
	s := server.NewServer()
	s.StartServer()
}

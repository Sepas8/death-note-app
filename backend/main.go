package main

import (
	"backend/server"
)

func main() {
	s := server.NewServer()
	s.StartServer()
}

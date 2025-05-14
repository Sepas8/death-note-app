package main

import (
	"github.com/Sepas8/death-note-app/backend/server"
)

func main() {
	s := server.NewServer()
	s.StartServer()
}

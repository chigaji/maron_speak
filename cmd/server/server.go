package main

import "github.com/chigaji/maron_speak/internal/server"

func main() {
	e := server.NewServer()
	e.Logger.Fatal(e.Start(":8080"))
}

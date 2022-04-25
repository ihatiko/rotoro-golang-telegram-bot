package main

import "rotoro-golang-telegram-bot/cmd"

func main() {
	server := cmd.NewServer()
	server.Serve()
}

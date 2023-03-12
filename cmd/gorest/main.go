package main

import (
	"github.com/captainherolee/clean-arch-go/internal/app/server"
)

// noinspection GoUnhandledErrorResult
func main() {
	server.NewServer().Start()
}

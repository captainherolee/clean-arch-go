package server

import (
	"github.com/captainherolee/clean-arch-go/internal/app/api"

	"github.com/captainherolee/go-rest/core"
)

func NewServer() core.Server {
	agApp := api.CreateAPIApp()
	server := core.NewServer(agApp, &(miscApp{}))
	_ = server.Init()
	return server
}

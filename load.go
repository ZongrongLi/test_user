package main

import (
	"github.com/tiancai110a/go-rpc/server"
	"github.com/tiancai110a/test_user/service/user"
)

// Load rpc service
func Load(s server.RPCServer) {
	s.Register(user.User{})
	return
}

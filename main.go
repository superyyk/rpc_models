package main

import (
	"flag"

	"github.com/smallnest/rpcx/server"
	baogai "github.com/superyyk/baogai_model"
)

var (
	addr = flag.String("addr", "localhost:39100", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.RegisterName("Arith", new(example.Arith), "")
	s.RegisterName("User", new(baogai.User), "")
	s.Serve("tcp", *addr)

}

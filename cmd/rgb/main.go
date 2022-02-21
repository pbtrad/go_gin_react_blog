package main

import (
	"go_gin_react_blog/internal/cli"
	"go_gin_react_blog/internal/conf"
	"go_gin_react_blog/internal/server"
)

func main() {
	env := cli.Parse()
	server.Start(conf.NewConfig(env))
}

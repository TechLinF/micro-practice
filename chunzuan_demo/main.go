package main

import (
	"git.rd.rijin.com/chunzuan_demo/handler"
	"git.rd.rijin.com/chunzuan_demo/infrastructure"
	"git.rd.rijin.com/chunzuan_demo/proto"
	"git.rd.rijin.com/micro/utils/server"
	// "github.com/gogap-micro/post-api/api/helper"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := infrastructure.Init()
	if err != nil {
		return
	}

	todoComponent := &handler.Todo{}

	srv := server.NewServer(
		server.Name(infrastructure.SERVICE))

	service_todo.RegisterTodoHandler(srv.Server(), todoComponent)

	err = srv.Run()
	if err != nil {
		return
	}
}

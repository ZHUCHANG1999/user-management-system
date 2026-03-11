package main

import (
	"flag"
	"fmt"
	"net/http"

	"user-management-system/internal/handler"
	"user-management-system/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "config/user.yaml", "the config file")

func main() {
	flag.Parse()

	ctx := svc.NewServiceContext()
	server := rest.MustNewServer(rest.RestConf{
		Host: "0.0.0.0",
		Port: 8888,
	})
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Println("Starting server at http://0.0.0.0:8888...")
	server.Start()
}

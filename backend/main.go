package main

import (
	"crud-sample/apps/auth"
	"github.com/rodrigorodriguescosta/goapp"
	//"github.com/rodrigorodriguescosta/goapp/adapters/grpc"
	"github.com/rodrigorodriguescosta/goapp/adapters/rest/fiber"
)

func main() {
	//grpcAdapter := grpc.New("1559")
	//grpcAdapter.Add(auth.App)
	restAdapter := fiber.New("8000")
	restAdapter.Add(auth.App)

	//goapp.AddAdapters(grpcAdapter, restAdapter)
	goapp.AddAdapters(restAdapter)
	goapp.Start("crud-sample")
}

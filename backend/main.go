package main

import (
	"github.com/waykiss/wkgo"
	"github.com/waykiss/wkproject/apps/auth"
	//"github.com/waykiss/wkgo/adapters/grpc"
	"github.com/waykiss/wkgo/adapters/rest/fiber"
)

func main() {
	//grpcAdapter := grpc.New("1559")
	//grpcAdapter.Add(auth.App)
	restAdapter := fiber.New("8000")
	restAdapter.Add(auth.App)

	//wkgo.AddAdapters(grpcAdapter, restAdapter)
	wkgo.AddAdapters(restAdapter)
	wkgo.Start("crud-sample")
}

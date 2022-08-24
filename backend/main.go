package main

import (
	"crud-sample/apps/auth"
	"goapp"
	//"goapp/adapters/grpc"
	"goapp/adapters/rest/fiber"
)

func main() {
	//grpcAdapter := grpc.New("1559")
	//grpcAdapter.Add(auth.App)

	restAdapter := fiber.New("6555")
	restAdapter.Add(auth.App)

	//goapp.AddAdapters(grpcAdapter, restAdapter)
	goapp.AddAdapters(restAdapter)
	goapp.Start("crud-sample")
}

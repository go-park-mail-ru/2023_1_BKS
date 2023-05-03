package main

import (
	"context"
	"fmt"
	"log"

	servGrpc "pkg/grpc/auth"

	"google.golang.org/grpc"
)

func main() {

	grcpConn, err := grpc.Dial(
		"127.0.0.1:8085",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	sessManager := servGrpc.NewAuthClient(grcpConn)

	cr := servGrpc.Id{Id: "243424"}

	wd, err := sessManager.GenerateAccessToken(context.Background(), &cr)

	fmt.Println(wd.GetValue())

}

package main

import (
	"context"
	"fmt"
	"log"

	servGrpc "pkg/grpc/user"

	"google.golang.org/grpc"
)

func main() {

	grcpConn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	sessManager := servGrpc.NewUserClient(grcpConn)

	cr := servGrpc.UserCheck{Login: "dw3234dw", Password: "5445"}

	wd, err := sessManager.CheckAccount(context.Background(), &cr)
	fmt.Println(wd.GetValue())

}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ninja-dark/test-assigment/grpcService"
	playclient "github.com/ninja-dark/test-assigment/internal/handler/client/playClient"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	client := playclient.NewPlayClient(conn)

	song, err := client.Player(context.TODO(), grpcService.PlayerRequest_play)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(song)

}

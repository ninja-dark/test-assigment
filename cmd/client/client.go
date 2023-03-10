package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit"
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

func AddSong(n int, client *playclient.PlayClient ){
	type song struct {
		a int64
		t string
		d time.Duration
	}

	songs := []song{}
	for i := n; i > 0; i-- {
		songs = append(songs, song{gofakeit.Int64(), gofakeit.Name(), time.Second * time.Duration(rand.Intn(150))})
	}
	for _, s := range songs {
		_, err := client.AddSong(
			context.TODO(),
			
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}
package main

import (
	"context"
	"github.com/brianvoe/gofakeit"
	"github.com/ninja-dark/test-assigment/internal/entity"
	playclient "github.com/ninja-dark/test-assigment/internal/handler/client/playClient"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	client := playclient.NewPlayClient(conn)

	AddSong(1, client)
	// вызов метода Play

	playResp, err := client.Play(context.Background())
	if err != nil {
		log.Fatalf("failed to Play: %v", err)
	}
	log.Printf("Play response: %t", playResp)

	// вызов метода Pause

	pauseResp, err := client.Pause(context.Background())
	if err != nil {
		log.Fatalf("failed to Pause: %v", err)
	}
	log.Printf("Pause response: %v", pauseResp)

	// вызов метода Next
	/*nextResp, nameNext, err := client.Next(context.Background())
	if err != nil {
		log.Fatalf("failed to Next: %v", err)
	}
	log.Printf("Next response: %t%v", nextResp, nameNext)

	// вызов метода Previous

	previousResp, namePrev, err := client.Previous(context.Background())
	if err != nil {
		log.Fatalf("failed to Previous: %v", err)
	}
	log.Printf("Previous response: %t%v", previousResp, namePrev)
	*/

}

func AddSong(n int, client *playclient.PlayClient) {
	songs := []entity.Song{}
	for i := n; i > 0; i-- {
		songs = append(songs, entity.Song{ID: gofakeit.Int64(), Name: gofakeit.Name(), Duration: time.Duration(gofakeit.Second())})
	}
	for _, s := range songs {
		_, err := client.AddSong(
			context.TODO(),
			&entity.Song{
				ID:       s.ID,
				Name:     s.Name,
				Duration: s.Duration,
			},
		)
		if err != nil {
			log.Print(err)
		}
	}
}

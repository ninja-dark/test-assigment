package playclient

import (
	"context"
	"fmt"

	"github.com/ninja-dark/test-assigment/grpcService"
	"github.com/ninja-dark/test-assigment/internal/entity"
	"google.golang.org/grpc"
)

type PlayClient struct {
	serviceCl grpcService.PlayCaseServicClient
}

func NewPlayClient(c *grpc.ClientConn) *PlayClient {
	return &PlayClient{
		serviceCl: grpcService.NewPlayCaseServicClient(c),
	}
}

func (c *PlayClient) AddSong(ctx context.Context, s *entity.Song) (int64, error) {
	r := &grpcService.AddSongRequest{Song: &grpcService.Song{
		Name:     s.Name,
		Duration: int64(s.Duration),
	},
	}
	song, err := c.serviceCl.AddSong(ctx, r)
	if err != nil {
		return 0, fmt.Errorf("can't add song: %w", err)
	}
	id := song.GetId()
	return id, nil
}

func (c *PlayClient) ReadSong(ctx context.Context) error {
	return nil
}

func (c *PlayClient) UpdateSong(ctx context.Context, s *entity.Song) (int64, error) {
	r := &grpcService.UpdateSongRequest{
		Song: &grpcService.Song{
			Id:       s.ID,
			Name:     s.Name,
			Duration: int64(s.Duration),
		},
	}
	song, err := c.serviceCl.UpdateSong(ctx, r)
	if err != nil {
		return 0, fmt.Errorf("can't update song: %w", err)
	}
	id := song.GetId()
	return id, nil
}

func (c *PlayClient) DeleteSong(ctx context.Context, id int64) (int64, error) {
	r := &grpcService.DeleteSongRequest{
		Id: id,
	}
	song, err := c.serviceCl.DeleteSong(ctx, r)
	if err != nil {
		return 0, fmt.Errorf("can't update song: %w", err)
	}
	deleteId := song.GetId()
	return deleteId, nil
}

func (c *PlayClient) Player(ctx context.Context, status grpcService.PlayerRequest_StatusPlayer) (string, error) {
	r := &grpcService.PlayerRequest{
		Status: status,
	}
	song, err := c.serviceCl.Player(ctx, r)
	if err != nil {
		return "", fmt.Errorf("can't change the status of a song: %w", err)
	}
	return song.Status, nil
}

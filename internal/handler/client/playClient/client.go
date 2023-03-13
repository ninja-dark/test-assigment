package playclient

import (
	"context"
	"fmt"
	"time"

	"github.com/ninja-dark/test-assigment/grpcService"
	"github.com/ninja-dark/test-assigment/internal/entity"
	"google.golang.org/grpc"
)

type PlayClient struct {
	serviceCl grpcService.MusicPlaylistClient
}

func NewPlayClient(c *grpc.ClientConn) *PlayClient {
	return &PlayClient{
		serviceCl: grpcService.NewMusicPlaylistClient(c),
	}
}

func (c *PlayClient) AddSong(ctx context.Context, s *entity.Song) (bool, error) {
	r := &grpcService.AddSongRequest{Song: &grpcService.Song{
		Title:    s.Name,
		Duration: int64(s.Duration),
	},
	}
	b, err := c.serviceCl.AddSong(ctx, r)
	if err != nil {
		return false, fmt.Errorf("can't add song: %w", err)
	}
	res := b.Success
	return res, nil
}

func (c *PlayClient) ReadSongs(ctx context.Context) ([]*entity.Song, error) {
	r := &grpcService.GetSongsRequest{}
	res, err := c.serviceCl.GetSongs(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("can't read song: %w", err)
	}
	sl := res.Song
	s := make([]*entity.Song, 0, len(sl))
	for _, song := range sl {
		s = append(s, &entity.Song{Name: song.Title, Duration: time.Duration(song.Duration)})
	}
	return s, nil
}

func (c *PlayClient) UpdateSong(ctx context.Context, s *entity.Song) (int64, error) {
	r := &grpcService.UpdateSongRequest{
		Song: &grpcService.Song{
			Id:       s.ID,
			Title:    s.Name,
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

func (c *PlayClient) DeleteSong(ctx context.Context, id int64) (bool, error) {
	r := &grpcService.DeleteSongRequest{
		Id: id,
	}
	song, err := c.serviceCl.DeleteSong(ctx, r)
	if err != nil {
		return false, fmt.Errorf("can't update song: %w", err)
	}

	return song.Success, nil
}

func (c *PlayClient) Play(ctx context.Context) (bool, error) {
	r := &grpcService.PlayRequest{}
	song, err := c.serviceCl.Play(ctx, r)
	if err != nil {
		return false, fmt.Errorf("can't change the status of a song: %w", err)
	}
	return song.Success, nil
}

func (c *PlayClient) Pause(ctx context.Context) (bool, error) {
	r := &grpcService.PauseRequest{}
	song, err := c.serviceCl.Pause(ctx, r)
	if err != nil {
		return false, fmt.Errorf("can't change the status of a song: %w", err)
	}
	return song.Success, nil
}

func (c *PlayClient) Next(ctx context.Context) (bool, string, error) {
	r := &grpcService.NextRequest{}
	song, err := c.serviceCl.Next(ctx, r)
	if err != nil {
		return false, "", fmt.Errorf("can't play next  of a song: %w", err)
	}
	return song.Success, song.Name, nil
}

func (c *PlayClient) Previous(ctx context.Context) (bool, string, error) {
	r := &grpcService.PreviousRequest{}
	song, err := c.serviceCl.Previous(ctx, r)
	if err != nil {
		return false, "", fmt.Errorf("can't play previous  of a song: %w", err)
	}
	return song.Success, song.Name, nil
}

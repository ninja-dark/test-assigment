package playcase

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ninja-dark/test-assigment/grpcService"
	"github.com/ninja-dark/test-assigment/internal/entity"
	"github.com/ninja-dark/test-assigment/internal/playlist"
)

type servic struct {
	playlist playlist.Playlist
	repo     entity.Repository
	grpcService.UnimplementedPlayCaseServicServer
}

func NewServic(playlist playlist.Playlist, repo entity.Repository) *servic {
	return &servic{
		playlist: playlist,
		repo:     repo,
	}
}

func (s *servic) AddSong(ctx context.Context, request *grpcService.AddSongRequest) (*grpcService.AddSongResponse, error) {
	err := ctx.Err()
	if err != nil {
		return nil, fmt.Errorf("context error: %w", err)
	}
	t := request.GetSong()
	song := &entity.Song{
		Name:     t.Name,
		Duration: time.Duration(t.Duration),
	}
	songPlaylist := playlist.Song{
		Name:     song.Name,
		Duration: song.Duration,
	}
	if err := s.playlist.AddSong(ctx, &songPlaylist); err != nil {
		return nil, fmt.Errorf("add song to playlist: %w", err)
	}
	p, err := s.repo.Add(ctx, song)
	if err != nil {
		return nil, fmt.Errorf("add song to database: %w", err)
	}

	return &grpcService.AddSongResponse{
		Id: p.ID,
	}, nil
}

func (s *servic) ReadSong(ctx context.Context, e *empty.Empty) (*grpcService.ReadSongResponse, error) {
	err := ctx.Err()
	if err != nil {
		return nil, fmt.Errorf("context error: %w", err)
	}
	t, err := s.repo.GetList(ctx)
	if err != nil {
		return nil, fmt.Errorf("get song: %w", err)
	}
	res := make([]*grpcService.Song, 0, len(t))
	for _, song := range t {
		res = append(res, &grpcService.Song{Name: song.Name, Duration: int64(song.Duration)})
	}
	return &grpcService.ReadSongResponse{Song: res}, nil
}

func (s *servic) UpdateSong(ctx context.Context, request *grpcService.UpdateSongRequest) (*grpcService.UpdateSongResponse, error) {
	err := ctx.Err()
	if err != nil {
		return nil, fmt.Errorf("context error: %w", err)
	}
	id := request.Song.GetId()
	name := request.Song.GetName()
	duration := request.Song.GetDuration()
	if id == 0 {
		return nil, fmt.Errorf("id not found")
	} else if name == "" && duration == 0 {
		return nil, fmt.Errorf("no data for update")
	}
	song := &entity.Song{
		ID:       id,
		Name:     name,
		Duration: time.Duration(duration),
	}
	idSong, err := s.repo.Update(ctx, song)
	if err != nil {
		return nil, fmt.Errorf("update song: %w", err)
	}
	return &grpcService.UpdateSongResponse{
		Id: idSong,
	}, nil
}

func (s *servic) DeleteSong(ctx context.Context, request *grpcService.DeleteSongRequest) (*grpcService.DeleteSongResponse, error) {
	err := ctx.Err()
	if err != nil {
		return nil, fmt.Errorf("context error: %w", err)
	}
	id := request.GetId()
	if id == 0 {
		return nil, fmt.Errorf("id not found")
	}
	if err := s.repo.Delete(ctx, id); err != nil {
		return nil, fmt.Errorf("delete song: %w", err)
	}
	return &grpcService.DeleteSongResponse{
		Id: id,
	}, nil
}

func (s *servic) Player(ctx context.Context, request *grpcService.PlayerRequest) (*grpcService.Status, error) {
	err := ctx.Err()
	if err != nil {
		return nil, fmt.Errorf("context error: %w", err)
	}
	var song *playlist.Song
	st := "action completed"
	status := request.GetStatus()
	switch status {
	case grpcService.PlayerRequest_play:
		song, err = s.playlist.Play(ctx)
	case grpcService.PlayerRequest_pause:
		song, err = s.playlist.Pause(ctx)
	case grpcService.PlayerRequest_next:
		song, err = s.playlist.Next(ctx)
	case grpcService.PlayerRequest_prev:
		song, err = s.playlist.Prev(ctx)
	}
	if err != nil {
		return nil, fmt.Errorf("delete song: %w", err)
	}
	return &grpcService.Status{
		Status: st,
		Name:   song.Name,
	}, nil
}

package playcase

import (
	"context"
	"fmt"
	"time"

	"github.com/ninja-dark/test-assigment/grpcService"
	"github.com/ninja-dark/test-assigment/internal/entity"
	"github.com/ninja-dark/test-assigment/internal/playlist"
)

type servic struct {
	playlist playlist.Playlist
	repo     entity.Repository
	grpcService.UnimplementedMusicPlaylistServer
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
		return &grpcService.AddSongResponse{
			Success: false,
		}, fmt.Errorf("context error: %w", err)
	}
	t := request.GetSong()
	song := &entity.Song{
		Name:     t.Title,
		Duration: time.Duration(t.Duration),
	}
	songPlaylist := playlist.Song{
		Name:     t.Title,
		Duration: song.Duration,
	}
	track := s.playlist.AddSong(ctx, &songPlaylist) 
	fmt.Printf(track.Name)
	err = s.repo.Add(ctx, song)
	if err != nil {
		return &grpcService.AddSongResponse{
			Success: false,
		}, fmt.Errorf("add song to database: %w", err)
	}

	return &grpcService.AddSongResponse{
		Success: true,
	}, nil
}

func (s *servic) GetSongs(ctx context.Context, request *grpcService.GetSongsRequest) (*grpcService.GetSongsResponse, error) {
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
		res = append(res, &grpcService.Song{Title: song.Name, Duration: int64(song.Duration)})
	}
	return &grpcService.GetSongsResponse{Song: res}, nil
}

func (s *servic) UpdateSong(ctx context.Context, request *grpcService.UpdateSongRequest) (*grpcService.UpdateSongResponse, error) {
	err := ctx.Err()
	if err != nil {
		return &grpcService.UpdateSongResponse{
			Success: false,
		}, fmt.Errorf("context error: %w", err)
	}
	id := request.Song.GetId()
	name := request.Song.GetTitle()
	duration := request.Song.GetDuration()
	if id == 0 {
		return &grpcService.UpdateSongResponse{
			Success: false,
		}, fmt.Errorf("id not found")
	} else if name == "" && duration == 0 {
		return &grpcService.UpdateSongResponse{
			Success: false,
		}, fmt.Errorf("no data for update")
	}
	song := &entity.Song{
		ID:       id,
		Name:     name,
		Duration: time.Duration(duration),
	}
	idSong, err := s.repo.Update(ctx, song)
	if err != nil {
		return &grpcService.UpdateSongResponse{
			Success: false,
		}, fmt.Errorf("update song: %w", err)
	}
	return &grpcService.UpdateSongResponse{
		Success: true,
		Id:      idSong,
	}, nil
}

func (s *servic) DeleteSong(ctx context.Context, request *grpcService.DeleteSongRequest) (*grpcService.DeleteSongResponse, error) {
	err := ctx.Err()
	if err != nil {
		return &grpcService.DeleteSongResponse{
			Success: false,
		}, fmt.Errorf("context error: %w", err)
	}
	id := request.GetId()
	if id == 0 {
		return &grpcService.DeleteSongResponse{
			Success: false,
		}, fmt.Errorf("id not found")
	}
	if err := s.repo.Delete(ctx, id); err != nil {
		return &grpcService.DeleteSongResponse{
			Success: false,
		}, fmt.Errorf("delete song: %w", err)
	}
	return &grpcService.DeleteSongResponse{
		Success: true,
	}, nil
}

func (s *servic) Play(ctx context.Context, request *grpcService.PlayRequest) (*grpcService.PlayResponse, error) {
	err := ctx.Err()
	if err != nil {
		return &grpcService.PlayResponse{
			Success: false,
		}, fmt.Errorf("context error: %w", err)
	}

	song, err := s.playlist.Play(ctx)
	if err != nil {
		return &grpcService.PlayResponse{
			Success: false,
		}, fmt.Errorf("play error: %w", err)
	}
	name := song.Name
	return &grpcService.PlayResponse{
		Success: true,
		Name:    name,
	}, nil
}

func (s *servic) Pause(ctx context.Context, request *grpcService.PauseRequest) (*grpcService.PauseResponse, error) {
	err := ctx.Err()
	if err != nil {
		return &grpcService.PauseResponse{
			Success: false,
		}, fmt.Errorf("context error: %w", err)
	}
	err = s.playlist.Pause(ctx)
	if err != nil {
		return &grpcService.PauseResponse{
			Success: false,
		}, fmt.Errorf("pause error: %w", err)
	}
	return &grpcService.PauseResponse{Success: true}, nil
}

func (s *servic) Next(ctx context.Context, request *grpcService.NextRequest) (*grpcService.NextResponse, error) {
	err := ctx.Err()
	if err != nil {
		return &grpcService.NextResponse{
			Success: false,
		}, fmt.Errorf("context error: %w", err)
	}
	song, err := s.playlist.Next(ctx)
	if err != nil {
		return &grpcService.NextResponse{
			Success: false,
		}, fmt.Errorf("next song error: %w", err)
	}
	
	return &grpcService.NextResponse{
		Success: true,
		Name:    song.Name,
	}, nil
}

func (s *servic) Previous(ctx context.Context, request *grpcService.PreviousRequest) (*grpcService.PreviousResponse, error) {
	err := ctx.Err()
	if err != nil {
		return &grpcService.PreviousResponse{
			Success: false,
		}, fmt.Errorf("context error: %w", err)
	}
	song, err := s.playlist.Prev(ctx)
	return &grpcService.PreviousResponse{
		Success: true,
		Name:    song.Name,
	}, nil
}

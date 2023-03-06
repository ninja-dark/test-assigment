package playcase

import (
	"context"
	"fmt"

	"github.com/ninja-dark/test-assigment/internal/entity"
	"github.com/ninja-dark/test-assigment/internal/playlist"
)

type servic struct {
	playlist playlist.Playlist
	repo     entity.Repository
}

type Servic interface {
	Play(ctx context.Context) (entity.Song ,error)
	Pause(ctx context.Context) error
	AddSong(ctx context.Context, song *entity.Song) (*entity.Song, error)
	NextSong(ctx context.Context) error
	PreviouSong(ctx context.Context) error
	UpdateSong(ctx context.Context) error
	DeleteSong(ctx context.Context) error
}

func NewServic(playlist playlist.Playlist, repo entity.Repository) Servic {
	return &servic{
		playlist: playlist,
		repo:     repo,
	}
}

func (s *servic) Play(ctx context.Context) (entity.Song ,error) {
	return s.playlist.Play(ctx, )
}

func (s *servic) Pause(ctx context.Context) error {
	return nil
}

func (s *servic) AddSong(ctx context.Context, song *entity.Song) (*entity.Song,error) {
	t := playlist.Song{
		Name: song.Name,
		Duration: song.Duration,
	}
	if err := s.playlist.AddSong(ctx, t); err != nil {
		return &entity.Song{},fmt.Errorf("add song to playlist: %w", err)
   }
   p,  err := s.repo.Add(ctx, song)
  	 if  err != nil {
		return &entity.Song{}, fmt.Errorf("add song to database: %w", err)
	}	
   
   return p, nil
}

func (s *servic) NextSong(ctx context.Context) error{
	return nil
}

func (s *servic) PreviouSong(ctx context.Context) error{
	return nil
}

func (s *servic) UpdateSong(ctx context.Context) error{
	return nil
}

func (s *servic) DeleteSong(ctx context.Context) error{
	return nil
}
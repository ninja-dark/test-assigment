package playlist

import (
	"context"
	"errors"
	"time"
)

type Song struct {
	ID       int64
	Name     string
	Duration time.Duration
	Status   string
}

type playlist struct {
	Playlist
	storage storage
}

type storage interface {
	Firstrack(ctx context.Context) (*Song, error)
	PushBack(ctx context.Context, song Song) error
	NextSong(ctx context.Context) error
	PrevSong(ctx context.Context, s Song) error
}

type Playlist interface {
	Play(ctx context.Context) (*Song, error)
	Pause(ctx context.Context) error
	AddSong(ctx context.Context, s Song) error
	Next(ctx context.Context) error
	Prev(ctx context.Context, s Song) error
}

func NewPlaylust() Playlist {
	return &playlist{}
}

func (p *playlist) Play(ctx context.Context) (*Song, error) {
	for track, err := p.storage.Firstrack(ctx); err != nil; err = p.Next(ctx) {
		d.currntTrack
		return track, nil
	}
	return &Song{}, nil
}

func (p *playlist) Pause(ctx context.Context) error {
	return nil
}
func (p *playlist) AddSong(ctx context.Context, s Song) error {
	err := p.storage.PushBack(ctx, s)
	if err != nil {
		return errors.New("can't add song")
	}
	return nil
}
func (p *playlist) Next(ctx context.Context) error {

	return nil
}
func (p *playlist) Prev(ctx context.Context, s Song) error {
	err := p.storage.PrevSong(ctx, s)
	if err != nil {
		return errors.New("can't find previous song")
	}
	return nil
}

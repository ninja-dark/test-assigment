package playlist

import (
	"context"
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
	PushBack()
	NextSong()
}

type Playlist interface {
	Play(ctx context.Context) error
	Pause(ctx context.Context) error
	AddSong(ctx context.Context, s Song) error
	Next(ctx context.Context) error
	Prev(ctx context.Context) error
}

func NewPlaylust() Playlist {
	return newDoubLink()
}

func (p *playlist) Play(ctx context.Context) error {
	return nil
}
func (p *playlist) Pause(ctx context.Context) error {
	return nil
}
func (p *playlist) AddSong(ctx context.Context, s Song) error {
	return nil
}
func (p *playlist) Next(ctx context.Context) error {
	return nil
}
func (p *playlist) Prev(ctx context.Context) error {
	return nil
}

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
}

type playlist struct {
	Playlist
	storage storage
}

type storage interface {
	Firstrack(ctx context.Context) (*Song, error)
	PushBack(ctx context.Context, song Song) *Song
	NextSong(ctx context.Context) (*Song, error)
	PrevSong(ctx context.Context) error
}

type Playlist interface {
	Play(ctx context.Context) (*Song, error)
	Pause(ctx context.Context) error
	AddSong(ctx context.Context, s Song) *Song
	Next(ctx context.Context) (*Song, error)
	Prev(ctx context.Context) error
}

func NewPlaylust() Playlist {
	return &playlist{
		storage: newDoubleLinkedList(),
	}
}

// Play начинает воспроизведение
func (p *playlist) Play(ctx context.Context) (*Song, error) {
	track, err := p.storage.Firstrack(ctx)
	if err != nil {
		return &Song{}, errors.New("playlist is empty")
	}
	track, err = p.storage.NextSong(ctx)
	if err != nil {
		return &Song{}, errors.New("playlist is empty")
	}
	return track, nil
}

// Pause приостанавливает воспроизведение
func (p *playlist) Pause(ctx context.Context) error {
	return nil
}

// AddSong добавляет в конец плейлиста песню
func (p *playlist) AddSong(ctx context.Context, s Song) *Song {
	track := p.storage.PushBack(ctx, s)
	return track
}

// Next воспроизвести след песню
func (p *playlist) Next(ctx context.Context) (*Song, error) {
	track, err := p.storage.NextSong(ctx)
	if err != nil {
		return &Song{}, errors.New("playlist is empty")
	}
	return track, nil
}

// Prev воспроизвести предыдущую песню
func (p *playlist) Prev(ctx context.Context) error {
	err := p.storage.PrevSong(ctx)
	if err != nil {
		return errors.New("can't find previous song")
	}
	return nil
}

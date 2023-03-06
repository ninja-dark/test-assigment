package playlist

import (
	"context"
	"time"
)

type playlistStatus int

const (
	playlistStatusStopped playlistStatus = iota
	playlistStatusPlaying
	playlistStatusPaused
)

type Song struct {
	ID       int64
	Name     string
	Duration time.Duration
}

type playlist struct {
	Playlist
	storage storage
	status  playlistStatus
}

type storage interface {
	Firstrack(ctx context.Context) (*Song, error)
	PushBack(ctx context.Context, song Song) *Song
	NextSong(ctx context.Context) (*Song, error)
	PrevSong(ctx context.Context) (*Song, error)
}

type Playlist interface {
	Play(ctx context.Context, s Song) (*Song, error)
	Pause(ctx context.Context, s Song) (*Song, error)
	AddSong(ctx context.Context, s Song) *Song
	Next(ctx context.Context) (*Song, error)
	Prev(ctx context.Context)  (*Song, error)
}

func NewPlaylist() Playlist {
	return &playlist{
		storage: newDoubleLinkedList(),
		status:  playlistStatusStopped,
	}
}

// Play начинает воспроизведение
func (p *playlist) Play(ctx context.Context, s Song) (*Song, error) {

	switch p.status {
	case playlistStatusStopped:
		track, err := p.storage.Firstrack(ctx)
		if err != nil {
			return &Song{}, ErrorPlaylistIsEmpty
		}
		p.status = playlistStatusPlaying
		track, err = p.storage.NextSong(ctx)
		if err != nil {
			return track, ErrorPlaylistIsEmpty
		}
	case playlistStatusPaused:
		track, err := p.Play(ctx, s)
		if err != nil {
			return track, ErrorPlaylistIsEmpty
		}
	case playlistStatusPlaying:
		track, err := p.Play(ctx, s)
		if err != nil{
			return track, ErrorPlayPlaylust
		}
	}
	return &Song{}, nil
}

// Pause приостанавливает воспроизведение
func (p *playlist) Pause(ctx context.Context, s Song) (*Song, error) {
	switch p.status{
	case playlistStatusPlaying:
		p.status = playlistStatusPaused
		return &Song{
			Duration: s.Duration,
		}, nil
	case playlistStatusPaused:
		return &s, ErrorPausePlaylust
	case playlistStatusStopped:
		return &s, ErrorStopPlaylust
	}
	return &Song{}, nil
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
		return &Song{}, ErrorPlaylistIsEmpty
	}
	t, err := p.Play(ctx, *track)
	return t, nil
}

// Prev воспроизвести предыдущую песню
func (p *playlist) Prev(ctx context.Context)  (*Song, error) {
	track, err := p.storage.PrevSong(ctx)
	if err != nil {
		return &Song{}, ErrorFindPrevSong
	}
	t, err := p.Play(ctx, *track)
	return t, nil
}

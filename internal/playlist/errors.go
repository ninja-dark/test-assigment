package playlist

import "errors"

var (
	ErrorPlaylistIsEmpty = errors.New("playlist is empty")
	ErrorFindPrevSong    = errors.New("can't find previous song")
	ErrorFindNextSong    = errors.New("can't find next song")
	ErrorPausePlaylust   = errors.New("playlist already paused")
	ErrorStopPlaylust    = errors.New("playlist stopped")
	ErrorPlayPlaylust    = errors.New("playlist already playing")
)

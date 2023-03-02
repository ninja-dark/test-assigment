package playlist

import (
	"container/list"
	"context"
	"errors"
	"sync"
)

type doubleLinkedList struct {
	track       *list.List
	currntTrack *list.Element
	lock        sync.Mutex
	storage
}

func newDoubleLinkedList() storage {
	return &doubleLinkedList{}
}

func (d *doubleLinkedList) PushBack(ctx context.Context, song Song) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.track.PushBack(song)
	return nil
}

func (d *doubleLinkedList) NextSong(ctx context.Context) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if d.track.Len() == 0 {
		return errors.New("playlist is empty")
	}
	d.currntTrack = d.currntTrack.Next()

	return nil
}

func (d *doubleLinkedList) PrevSong(ctx context.Context) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.currntTrack = d.currntTrack.Prev()
	if d.currntTrack == nil {
		d.currntTrack = d.track.Back()
	}
	return nil
}

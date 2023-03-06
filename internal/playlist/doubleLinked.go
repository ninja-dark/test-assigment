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
	return &doubleLinkedList{
		track:       list.New(),
		currntTrack: nil,
	}
}

// Firstrack возвращает первый элемент из списка
func (d *doubleLinkedList) Firstrack(ctx context.Context) (*Song, error) {
	d.currntTrack = d.track.Front()
	if d.currntTrack == nil {
		return &Song{}, errors.New("playlist is empty")
	}
	track := d.currntTrack.Value.(*Song)
	return track, nil
}

// PushBack вставляет новый элемент в конец списка
func (d *doubleLinkedList) PushBack(ctx context.Context, song Song) *Song {
	d.lock.Lock()
	defer d.lock.Unlock()

	t := d.track.PushBack(song)
	track := t.Value.(*Song)
	return track
}

// NextSong возвращает следующий элемент из списка
func (d *doubleLinkedList) NextSong(ctx context.Context) (*Song, error) {
	d.lock.Lock()
	defer d.lock.Unlock()

	if d.track.Len() == 0 {
		return nil, errors.New("playlist is empty")
	}
	d.currntTrack = d.currntTrack.Next()
	if d.currntTrack == nil {
		d.currntTrack = d.track.Front()
	}
	track := d.currntTrack.Value.(*Song)

	return track, nil
}

// PrevSong возвращает предыдущий элемент из списка
func (d *doubleLinkedList) PrevSong(ctx context.Context) (*Song, error) {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.currntTrack = d.currntTrack.Prev()
	if d.currntTrack == nil {
		d.currntTrack = d.track.Back()
	}
	track := d.currntTrack.Value.(*Song)
	return track, nil
}

package entity

import (
	"context"
	"time"
)

type Song struct {
	ID       int64
	Name     string
	Duration time.Duration
}

type Repository interface {
	GetList(ctx context.Context) ([]*Song, error)
	Add(ctx context.Context, s *Song) error
	Update(ctx context.Context, s *Song) (int64, error)
	Delete(ctx context.Context, id int64) error
}

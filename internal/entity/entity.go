package entity

import "time"

type Song struct {
	ID       int64
	Name     string
	Duration time.Duration
	Status   string
}

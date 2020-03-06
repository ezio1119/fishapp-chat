package domain

import (
	"time"
)

type Room struct {
	ID        int64
	PostID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

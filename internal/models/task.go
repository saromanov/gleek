package models

import "time"

// Task defines task model
type Task struct {
	ID        int64
	CreatedAt time.Time     `db:"created_at"`
	Name      string        `db:"name"`
	Priority  uint          `db:"priority"`
	Start     time.Time     `db:"start"`
	Duration  time.Duration `db:"duration"`
}

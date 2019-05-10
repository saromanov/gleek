package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Task defines task model
type Task struct {
	gorm.Model
	Name     string
	Priority uint
	Start    time.Time
	Duration time.Duration
	Tags     []Tag `gorm:"many2many:task_tags;omitempty"`
}

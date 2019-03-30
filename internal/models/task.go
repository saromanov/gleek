package models

import (
	"github.com/jinzhu/gorm"
)

// Task defines task model
type Task struct {
	gorm.Model
	Name     string
	Priority uint
}

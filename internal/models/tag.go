package models

import (
	"github.com/jinzhu/gorm"
)

// Tag defines tag model
type Tag struct {
	gorm.Model
	Name string
}

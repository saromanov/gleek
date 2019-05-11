package storage

import (
	"github.com/saromanov/gleek/internal/models"
	pb "github.com/saromanov/gleek/proto"
)

// CreateTask provides inserting of the task
func (s *Storage) CreateTask(t *pb.Task) (uint, error) {
	return 0, nil
}

// GetTask provides provides getting of the task
func (s *Storage) GetTask(id uint) (*models.Task, error) {
	return nil, nil
}

package storage

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/saromanov/gleek/internal/models"
	pb "github.com/saromanov/gleek/proto"
)

// CreateTask provides inserting of the task
func (s *Storage) CreateTask(t *pb.Task) (uint, error) {
	task := &models.Task{
		Name: t.Name,
	}
	tags := make([]models.Tag, len(t.Tags))
	for i, tag := range t.Tags {
		tags[i] = models.Tag{
			Name: tag.Name,
		}
	}
	task.Tags = tags
	err := s.db.Create(task).Error
	if err != nil {
		return 0, errors.Wrap(err, "storage: unable to insert user")
	}
	return 0, nil
}

// GetTask provides provides getting of the task
func (s *Storage) GetTask(id uint) (*models.Task, error) {
	var tags []models.Tag
	task := &models.Task{}
	err := s.db.First(&task, "id = ?", id).Error
	if err != nil {
		panic(err)
	}
	s.db.Model(&task).Related(&tags, "Tags")
	fmt.Println(tags)
	fmt.Println(task)
	return task, nil
}

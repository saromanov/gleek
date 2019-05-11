package storage

import (
	"time"

	"github.com/pkg/errors"
	"github.com/saromanov/gleek/internal/models"
	pb "github.com/saromanov/gleek/proto"
)

// CreateTask provides inserting of the task
func (s *Storage) CreateTask(t *pb.Task) (uint, error) {
	result, err := s.db.Exec(taskSchema)
	if err != nil {
		return 0, errors.Wrap(err, "unable to execute task schema")
	}
	newTask := `INSERT INTO tasks (created_at, name, priority, start, duration) VALUES (?, ?, ?, ?, ?, ?)`
	s.db.MustExec(newTask, time.Now().UTC(), t.Name, t.Priority, t.Start, t.DUration).Result()
	return 0, nil
}

// GetTask provides provides getting of the task
func (s *Storage) GetTask(id uint) (*models.Task, error) {
	return nil, nil
}

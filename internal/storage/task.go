package storage

import (
	"time"

	"github.com/pkg/errors"
	"github.com/saromanov/gleek/internal/models"
	pb "github.com/saromanov/gleek/proto"
)

// CreateTask provides inserting of the task
func (s *Storage) CreateTask(t *pb.Task) (int64, error) {
	result, err := s.db.Exec(taskSchema)
	if err != nil {
		return 0, errors.Wrap(err, "unable to execute task schema")
	}
	newTask := `INSERT INTO tasks (created_at, name, priority, duration) VALUES (?, ?, ?, ?, ?, ?)`
	s.db.MustExec(newTask, time.Now().UTC(), t.Name, t.Priority, t.Duration)
	return result.LastInsertId()
}

// GetTask provides provides getting of the task
func (s *Storage) GetTask(id uint) (*models.Task, error) {
	task := models.Task{}
	err := s.db.Get(&task, "SELECT * FROM tasks WHERE id=? LIMIT 1", id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get task")
	}
	return &task, nil
}

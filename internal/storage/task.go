package storage

import (
	"fmt"

	//"github.com/jinzhu/gorm"
	//"github.com/jinzhu/gorm"
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
	err := s.db.Debug().Create(task).Error
	if err != nil {
		return 0, errors.Wrap(err, "storage: unable to insert user")
	}
	return 0, nil
}

// GetTask provides provides getting of the task
func (s *Storage) GetTask(id uint) (*models.Task, error) {
	//var tags []models.Tag
	tasks := []*models.Task{}
	/*err := s.db.First(&task, "id = ?", id).Error
	if err != nil {
		panic(err)
	}*/

	/*s.db.Debug().Preload("Tags", func(db *gorm.DB) *gorm.DB {
		return db.Select("name").Where("name = ? OR name = ?", "lang", "ci").Order("name DESC")
	}).Find(&tasks)*/
	s.db.Debug().Table("task_tags").Select("*").Joins("inner join tags on task_tags.tag_id = tags.id").Where("tags.name = ?", "lang").Find(&tasks)
	//s.db.Model(&task).Related(&tags, "Tags")
	for _, ta := range tasks {
		fmt.Println(ta)
	}
	return nil, nil
}

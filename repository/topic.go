package repository

import (
	"log"
	"sync"
	"time"
)

type Topic struct {
	Id         int64     `gorm:"column:id"`
	Title      string    `gorm:"column:title"`
	Content    string    `gorm:"column:content"`
	CreateTime time.Time `gorm:"column:create_time"`
}
type TopicDAO struct {
}

var (
	topicDAO  *TopicDAO
	topicOnce sync.Once
)

func NewTopicDaoInstance() *TopicDAO {
	topicOnce.Do(
		func() {
			topicDAO = &TopicDAO{}
		})
	return topicDAO
}

func (*TopicDAO) QueryTopicById(id int64) (*Topic, error) {
	var topic Topic
	err := db.Where("id = ?", id).Find(&topic).Error
	if err != nil {
		log.Println("find topic by id err:" + err.Error())
		return nil, err
	}
	return &topic, nil
}

func (*TopicDAO) AddNewTopic(topic *Topic) (int64, error) {
	topic.CreateTime = time.Now()
	result := db.Create(topic)
	err := result.Error
	if err != nil {
		log.Println("insert post err:" + err.Error())
	}
	return topic.Id, err
}

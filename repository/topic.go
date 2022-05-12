package repository

import (
	"sync"
)

type Topic struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
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

func (*TopicDAO) QueryTopicById(id int64) *Topic {
	return topicIndexMap[id]
}

func (*TopicDAO) AddNewTopic(topic *Topic) (int64, error) {
	topic.Id = 3
	return 3, nil
}

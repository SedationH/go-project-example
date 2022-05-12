package repository

import (
	"encoding/json"
	"os"
	"sync"
	"time"
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
	topic, _ := topicIndexMap.Get(id)
	return topic
}

func (*TopicDAO) AddNewTopic(topic *Topic) (int64, error) {
	maxTopicId++
	topic.Id = maxTopicId
	topic.CreateTime = time.Now().Unix()
	topicIndexMap.Set(topic.Id, topic)
	f, err := os.OpenFile("./data/topic", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return -1, err
	}
	defer f.Close()
	f.WriteString("\n")
	b, _ := json.Marshal(topic)
	_, err = f.Write(b)
	if err != nil {
		return -1, err
	}
	return maxTopicId, nil
}

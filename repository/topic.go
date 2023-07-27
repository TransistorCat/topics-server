package repository

import (
	"encoding/json"
	"os"
	"sync"
)

type Topic struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}
type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do( //当且仅当第一次为此 Once 实例调用 Do 时，Do 才会调用函数 f
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}
func (*TopicDao) QueryTopicByID(id int64) *Topic {
	return topicIndexMap[id]
}

func (*TopicDao) InsertTopic(topic *Topic) error {
	f, err := os.OpenFile("./data/topic", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()
	marshal, _ := json.Marshal(topic)
	if _, err = f.WriteString(string(marshal) + "\n"); err != nil {
		return err
	}

	rwMutex.Lock()
	topicIndexMap[topic.ID] = topic

	rwMutex.Unlock()
	return nil
}

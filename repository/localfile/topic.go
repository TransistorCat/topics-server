package localfile

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/TransistorCat/topics-server/repository"
)

type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

func (*TopicDao) QueryTopicByID(id int64) *repository.Topic {
	return topicIndexMap[id]
}

func (*TopicDao) InsertTopic2Local(topic *repository.Topic) error {
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

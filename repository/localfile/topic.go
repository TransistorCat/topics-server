package localfile

import (
	"encoding/json"
	"os"

	. "github.com/TransistorCat/topics-server/repository/common"
)

type TopicDao struct{}

func NewTopicDao() *TopicDao {
	return &TopicDao{}
}
func (*TopicDao) QueryByID(id int64) *Topic {
	return topicIndexMap[id]
}

func (*TopicDao) Insert(topic *Topic) error {
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

package common

import (
	"sync"
)

type Topic struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

func (Topic) TableName() string {
	return "topic"
}

type TopicDao interface {
	QueryTopicByID(id int64) *Topic
	InsertTopic2Local(topic *Topic) error
}

var topicOnce sync.Once

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do( //当且仅当第一次为此 Once 实例调用 Do 时，Do 才会调用函数 f
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}

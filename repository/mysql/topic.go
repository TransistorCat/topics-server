package mysql

import "github.com/TransistorCat/topics-server/repository"

type TopicDao struct{}

var topicDao *TopicDao

func (*TopicDao) QueryTopicByID(id int64) *repository.Topic {
	var topic repository.Topic
	DB.Where("id=?", id).Find(&topic)
	return &topic
}

func (*TopicDao) InsertTopic(topic *repository.Topic) error {
	if err := DB.Create(topic).Error; err != nil {
		return err
	}

	return nil
}

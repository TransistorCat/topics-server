package mysql

import . "github.com/TransistorCat/topics-server/repository/common"

type TopicDao struct{}

var topicDao *TopicDao

func NewTopicDao() *TopicDao {
	return &TopicDao{}
}
func (*TopicDao) QueryByID(id int64) *Topic {
	var topic Topic
	DB.Where("id=?", id).Find(&topic)
	return &topic
}

func (*TopicDao) Insert(topic *Topic) error {
	if err := DB.Create(topic).Error; err != nil {
		return err
	}

	return nil
}

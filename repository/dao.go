package repository

import (
	"sync"

	. "github.com/TransistorCat/topics-server/repository/common"
	"github.com/TransistorCat/topics-server/repository/localfile"
	"github.com/TransistorCat/topics-server/repository/mysql"
)

var postDao PostDao
var postOnce sync.Once

type PostDao interface {
	QueryByParentID(parentid int64) []*Post
	Insert(post *Post) error
}

func NewPostDaoInstance(dbType DBType) PostDao {
	postOnce.Do(func() {
		switch dbType {
		case 1:
			postDao = localfile.NewPostDao()
		case 2:
			postDao = mysql.NewPostDao()
		}
	})
	return postDao
}

type TopicDao interface {
	QueryByID(id int64) *Topic
	Insert(topic *Topic) error
}

var topicOnce sync.Once
var topicDao TopicDao

func NewTopicDaoInstance(dbType DBType) TopicDao {
	topicOnce.Do( //当且仅当第一次为此 Once 实例调用 Do 时，Do 才会调用函数 f
		func() {
			switch dbType {
			case 1:
				topicDao = localfile.NewTopicDao()
			case 2:
				topicDao = mysql.NewTopicDao()
			}
		})
	return topicDao
}

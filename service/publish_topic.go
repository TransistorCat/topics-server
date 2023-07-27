package service

import (
	"errors"
	"time"
	"unicode/utf16"

	"github.com/TransistorCat/topics-server/repository"
	idworker "github.com/gitstliu/go-id-worker"
)

var idGen *idworker.IdWorker

func init() {
	idGen = &idworker.IdWorker{}
	idGen.InitIdWorker(1, 1)
}

// PublishTopic 用于发布帖子的函数。
func PublishTopic(title string, content string) (int64, error) {
	return NewPublishTopicFlow(title, content).Do()
}

// NewPublishTopicFlow 创建一个新的发布帖子的工作流实例。
func NewPublishTopicFlow(title string, content string) *PublishTopicFlow {
	return &PublishTopicFlow{
		title:   title,
		content: content,
	}
}

// PublishTopicFlow 表示发布帖子的工作流程。
type PublishTopicFlow struct {
	title   string
	content string
	topicID int64
}

// Do 执行发布帖子的工作流程。
func (f *PublishTopicFlow) Do() (int64, error) {
	// 检查参数有效性
	if err := f.checkParam(); err != nil {
		return 0, err
	}
	// 发布帖子
	if err := f.publish(); err != nil {
		return 0, err
	}
	return f.topicID, nil
}

// checkParam 检查参数有效性，确保内容长度在限制范围内。
func (f *PublishTopicFlow) checkParam() error {
	if len(utf16.Encode([]rune(f.title))) >= 30 {
		return errors.New("标题长度必须小于30")
	}
	if len(utf16.Encode([]rune(f.content))) >= 500 {
		return errors.New("内容长度必须小于500")
	}
	return nil
}

// publish 发布帖子，将帖子内容保存到数据库中。
func (f *PublishTopicFlow) publish() error {
	// 创建待发布的帖子对象
	topic := &repository.Topic{
		Title:      f.title,
		Content:    f.content,
		CreateTime: time.Now().Unix(),
	}
	// 生成唯一的帖子ID
	id, err := idGen.NextId()
	if err != nil {
		return err
	}
	topic.ID = id
	// 将帖子插入数据库
	if err := repository.NewTopicDaoInstance().InsertTopic(topic); err != nil {
		return err
	}
	f.topicID = topic.ID
	return nil
}

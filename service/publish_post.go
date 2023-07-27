package service

import (
	"errors"
	"time"
	"unicode/utf16"

	"github.com/TransistorCat/topics-server/repository"
	idworker "github.com/gitstliu/go-id-worker"
)

var postIDGen *idworker.IdWorker

func init() {
	postIDGen = &idworker.IdWorker{}
	postIDGen.InitIdWorker(1, 1)
}

func PublishPost(parentID int64, content string) (int64, error) {
	return NewPublishPostFlow(parentID, content).Do()
}

func NewPublishPostFlow(parentID int64, content string) *PublishPostFlow {
	return &PublishPostFlow{
		parentID: parentID,
		content:  content,
	}
}

type PublishPostFlow struct {
	postID   int64
	parentID int64
	content  string
}

func (f *PublishPostFlow) Do() (int64, error) {
	// 检查参数有效性
	if err := f.checkParam(); err != nil {
		return 0, err
	}
	// 发布帖子
	if err := f.publish(); err != nil {
		return 0, err
	}
	return f.postID, nil
}

func (f *PublishPostFlow) checkParam() error {

	if len(utf16.Encode([]rune(f.content))) >= 500 {
		return errors.New("内容长度必须小于500")
	}
	return nil
}

func (f *PublishPostFlow) publish() error {
	post := &repository.Post{
		ParentID:   f.parentID,
		Content:    f.content,
		CreateTime: time.Now().Unix(),
	}
	id, err := postIDGen.NextId()
	if err != nil {
		return nil
	}

	post.ID = id
	if err := repository.NewPostDaoInstance().InsertPost2MySQL(post); err != nil {
		return err
	}
	f.postID = post.ID
	return nil

}

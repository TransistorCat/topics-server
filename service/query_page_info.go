package service

import (
	"errors"
	"sync"

	"github.com/TransistorCat/topics-server/repository"
	. "github.com/TransistorCat/topics-server/repository/common"
)

type PageInfo struct {
	Topic    *Topic
	PostList []*Post
}

var needUpdate bool

// 通过主题ID来获取页面的信息
func QueryPageInfo(topicID int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(topicID).Do()
}

func NewQueryPageInfoFlow(topID int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicID: topID,
	}
}

// 与数据的交互
type QueryPageInfoFlow struct {
	topicID  int64
	pageInfo *PageInfo //目标

	topic *Topic
	posts []*Post
}

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}

func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicID <= 0 {
		return errors.New("topic id must be larger than 0")
	}
	return nil
}

func (f *QueryPageInfoFlow) prepareInfo() error {
	//获取topic信息
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if topic := repository.QueryTopicCache(f.topicID); topic != nil {
			f.topic = topic
			return
		}
		topic := repository.NewTopicDaoInstance(repository.DefaultOptions.DBType).QueryByID(f.topicID)
		repository.AppendTopicCache(*topic)
		f.topic = topic
	}()
	//获取post列表
	go func() {
		defer wg.Done()
		if posts := repository.QuerypostCache(f.topicID); posts != nil {
			f.posts = posts
			return
		}
		posts := repository.NewPostDaoInstance(repository.DefaultOptions.DBType).QueryByParentID(f.topicID)
		repository.AppendpostCache(posts)
		f.posts = posts
	}()
	wg.Wait()
	return nil
}

func (f *QueryPageInfoFlow) packPageInfo() error {
	f.pageInfo = &PageInfo{
		Topic:    f.topic,
		PostList: f.posts,
	}
	return nil
}

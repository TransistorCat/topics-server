package localfile

import (
	"bufio"
	"encoding/json"
	"os"
	"sync"

	. "github.com/TransistorCat/topics-server/repository/common"
)

var (
	topicIndexMap map[int64]*Topic
	postIndexMap  map[int64][]*Post
	rwMutex       sync.RWMutex
)

type localFile struct {
	filePath string
}

var Default = localFile{filePath: "./data/"}

func Init(f *localFile) error {

	if err := f.initTopicIndexMap(); err != nil {
		return err
	}
	if err := f.initPostIndexMap(); err != nil {
		return err
	}
	return nil
}
func (f *localFile) initTopicIndexMap() error {
	open, err := os.Open(f.filePath + "topic")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.ID] = &topic
	}
	topicIndexMap = topicTmpMap
	return nil
}

func (f *localFile) initPostIndexMap() error {
	open, err := os.Open(f.filePath + "post")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}
		posts, ok := postTmpMap[post.ParentID]
		if !ok { //不存在就新建
			postTmpMap[post.ParentID] = []*Post{&post}
			continue
		}
		posts = append(posts, &post)
		postTmpMap[post.ParentID] = posts
	}
	postIndexMap = postTmpMap
	return nil
}

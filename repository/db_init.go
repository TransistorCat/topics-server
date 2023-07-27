package repository

import (
	"bufio"
	"encoding/json"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	topicIndexMap map[int64]*Topic
	postIndexMap  map[int64][]*Post
	rwMutex       sync.RWMutex
	DB            *gorm.DB
)

func Init(options Options) error {
	switch options.DBType {
	case 1:
		initIndexMapFormLocal(options.FilePath)
	case 2:
		dataSourceName := "root:root@tcp(127.0.0.1:3306)/TopicServerDB?charset=utf8&parseTime=True"
		var err error
		if DB, err = gorm.Open(mysql.Open(dataSourceName), nil); err != nil {
			return err
		}
		initPostIndexMapFormMySQL(DB)
		initTopicIndexMapFormMySQL(DB)

	}
	return nil
}

func initIndexMapFormLocal(filePath string) error {
	if err := initTopicIndexMapFormLocal(filePath); err != nil {
		return err
	}
	if err := initPostIndexMapFormLocal(filePath); err != nil {
		return err
	}
	return nil
}

func initTopicIndexMapFormLocal(filePath string) error {
	open, err := os.Open(filePath + "topic")
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

func initPostIndexMapFormLocal(filePath string) error {
	open, err := os.Open(filePath + "post")
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

func initTopicIndexMapFormMySQL(DB *gorm.DB) error {
	var topics []Topic
	if err := DB.Find(&topics).Error; err != nil {
		return err
	}
	topicTmpMap := make(map[int64]*Topic)
	for _, topic := range topics {
		topicTmpMap[topic.ID] = &topic
	}
	topicIndexMap = topicTmpMap
	return nil
}

func initPostIndexMapFormMySQL(DB *gorm.DB) error {
	var posts []Post
	if err := DB.Find(&posts).Error; err != nil {
		return err
	}
	postTmpMap := make(map[int64][]*Post)
	for _, post := range posts {
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

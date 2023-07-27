package repository

import (
	"encoding/json"
	"os"
	"sync"
)

type Post struct {
	ID         int64  `json:"id"`
	ParentID   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type PostDao struct{}

var (
	postdao  *PostDao
	postOnce sync.Once
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(func() { postdao = &PostDao{} })
	return postdao
}

func (*PostDao) QueryPostsByParentID(parentid int64) []*Post {
	return postIndexMap[parentid]
}

func (*PostDao) InsertPost(post *Post) error {
	f, err := os.OpenFile("./data/post", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	marshal, _ := json.Marshal(post)
	if _, err := f.WriteString(string(marshal) + "\n"); err != nil {
		return err
	}
	rwMutex.Lock()
	postList, ok := postIndexMap[post.ParentID]
	if !ok {
		postIndexMap[post.ParentID] = []*Post{post}
	} else {
		postList = append(postList, post)
		postIndexMap[post.ParentID] = postList
	}
	rwMutex.Unlock()
	return nil
}

package common

import (
	"sync"
)

type Post struct {
	ID         int64  `json:"id"`
	ParentID   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

func (Post) TableName() string {
	return "post"
}

// type PostDao struct{}

var postOnce sync.Once

type PostDao interface {
	QueryPostsByParentID(parentid int64)
	InsertPost(post *Post) error
}

func NewPostDaoInstance() *PostDao {
	postOnce.Do(func() { postdao = &PostDao{} })
	return postdao
}

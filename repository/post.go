package repository

import "sync"

type Post struct {
	ID         int64  `json:id`
	ParentID   int64  `json:parent_id`
	Content    string `josn:content`
	CreateTime int64  `json:create_time`
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

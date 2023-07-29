package localfile

import (
	"encoding/json"
	"os"

	. "github.com/TransistorCat/topics-server/repository/common"
)

type PostDao struct{}

func NewPostDao() *PostDao {
	return &PostDao{}
}

func (*PostDao) QueryByParentID(parentid int64) []*Post {
	return postIndexMap[parentid]
}

func (*PostDao) Insert(post *Post) error {
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

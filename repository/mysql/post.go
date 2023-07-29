package mysql

import . "github.com/TransistorCat/topics-server/repository/common"

type PostDao struct{}

var (
	postdao *PostDao
)

func NewPostDao() *PostDao {
	return &PostDao{}
}

func (*PostDao) QueryByParentID(parentid int64) []*Post {
	var posts []*Post
	DB.Where("parent_id=?", parentid).Find(&posts)
	return posts
}

func (*PostDao) Insert(post *Post) error {
	if err := DB.Create(post).Error; err != nil {
		return err
	}
	return nil
}

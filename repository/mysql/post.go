package mysql

import "github.com/TransistorCat/topics-server/repository"

type PostDao struct{}

var (
	postdao *PostDao
)

func (*PostDao) QueryPostsByParentID(parentid int64) []*repository.Post {
	var posts []*repository.Post
	DB.Where("parent_id=?", parentid).Find(&posts)
	return posts
}

func (*PostDao) InsertPost(post *repository.Post) error {
	if err := DB.Create(post).Error; err != nil {
		return err
	}
	return nil
}

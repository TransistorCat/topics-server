package common

type Post struct {
	ID         int64  `json:"id"`
	ParentID   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

func (Post) TableName() string {
	return "post"
}

package common

import "encoding/json"

type Post struct {
	ID         int64  `json:"id"`
	ParentID   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

func (Post) TableName() string {
	return "post"
}

func (s *Post) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Post) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

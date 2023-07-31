package common

import "encoding/json"

type Topic struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

func (Topic) TableName() string {
	return "topic"
}

func (s *Topic) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Topic) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

package cotroller

import "github.com/TransistorCat/topics-server/service"

func PublishTopic(title string, content string) *PageData {
	topicID, err := service.PublishTopic(title, content)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &PageData{
		Code: 0,
		Msg:  "success",
		Data: map[string]int64{
			"topic_id": topicID,
		},
	}
}

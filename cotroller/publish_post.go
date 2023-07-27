package cotroller

import "github.com/TransistorCat/topics-server/service"

func PublishPost(parentID int64, content string) *PageData {
	postID, err := service.PublishPost(parentID, content)
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
			"post_id": postID,
		},
	}
}

package repository

import (
	"context"
	"fmt"

	. "github.com/TransistorCat/topics-server/repository/common"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func redisInit() {

	rdb = redis.NewClient(&redis.Options{
		Addr:     "43.138.243.212:6390",
		Password: "rootredis",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错，错误信息：%v", err)
	}
}

func QueryTopicCache(id int64) *Topic {
	if isExist, _ := rdb.HExists(ctx, "topic", fmt.Sprint(id)).Result(); isExist == false {
		return nil
	}
	var topic Topic
	err := rdb.HGet(ctx, "topic", fmt.Sprint(id)).Scan(&topic)
	if err != nil {
		panic(err)
	}

	return &topic
}

func AppendTopicCache(topic Topic) {
	keys, err := rdb.HSet(ctx, "topic", topic.ID, &topic).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)
}

func QuerypostCache(parentID int64) []*Post {
	if isExist, _ := rdb.Exists(ctx, "post"+fmt.Sprint(parentID)).Result(); isExist == 0 {
		return nil
	}
	var post []*Post
	// post = make([]*Post, 64)
	err := rdb.HVals(ctx, "post"+fmt.Sprint(parentID)).ScanSlice(&post)
	if err != nil {
		panic(err)
	}

	return post
}

func AppendpostCache(post []*Post) {
	for i := 0; i < len(post); i++ {
		keys, err := rdb.HSet(ctx, "post"+fmt.Sprint(post[i].ParentID), post[i].ID, &post[i]).Result()
		if err != nil {
			panic(err)
		}
		fmt.Println(keys)
	}

}

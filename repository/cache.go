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

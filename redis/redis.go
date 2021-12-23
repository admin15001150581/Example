package redis

import (
	"context"
	"example/pkg/setting"
	"github.com/go-redis/redis/v8"
)

//全局redis对象
var rdb *redis.Client

var ctx context.Context

//链接Redis
func init(){
	rdb = redis.NewClient(&redis.Options{
	Addr:	  setting.Viper.GetString("redis.Addr"),
	Password: "", // no password set
	DB:		  0,  // use default DB
	})

	ctx = context.Background()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic("init redis error")

	}
}
package basic

import (
	"github.com/basic/redis"
	"github.com/basic/config"
	"github.com/basic/db"
	"github.com/basic/mongo"
)

func Init()  {
	config.Init()
	db.Init()
	redis.Init()
	mongo.Init()

}
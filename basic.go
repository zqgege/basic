package basic

import (
	"microProject/basic/redis"
	"microProject/basic/config"
	"microProject/basic/db"
	"microProject/basic/mongo"
)

func Init()  {
	config.Init()
	db.Init()
	redis.Init()
	mongo.Init()

}
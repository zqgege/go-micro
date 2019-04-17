package basic

import (
	"microProject/basic/db"
	"microProject/basic/config"
	"microProject/basic/redis"
)

func Init()  {
	config.Init()
	db.Init()
	redis.Init()

}
package mongo

import (
	"github.com/go-log/log"
	"gopkg.in/mgo.v2"
	"microProject/basic/config"
	_"strconv"
	"sync"
	"time"
)

var (
	mongoSession *mgo.Session
	Db *mgo.Database
	m sync.RWMutex
	inited bool
)

func Init()  {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Logf("已经初始化过Mongo...")
		return
	}
	mongoConfig := config.GetMongoConfig()
	// 打开才加载
	if mongoConfig != nil && mongoConfig.GetEnabled(){
		log.Log("初始化mongo...")
		/*if mongoSession == nil {
			var err error
			mongoSession, err = mgo.Dial(mongoConfig.GetConn())
			poolLimit, _    := strconv.Atoi(string(mongoConfig.GETPoolLimit()))
			mongoSession.SetPoolLimit(poolLimit)

			if err != nil {
				panic(err) // no, not really
			}
		}
		//mongoSession.DB(mongoConfig.GETDb())
		mongoSession.Clone()*/

		var err error
		dialInfo := &mgo.DialInfo{
			Addrs:     []string{mongoConfig.GetConn()},
			Direct:    false,
			Timeout:   time.Second * 1,
			PoolLimit: 4096,
		}
		//创建一个维护套接字池的session
		mongoSession, err = mgo.DialWithInfo(dialInfo)
		if err != nil {
			log.Log("创建一个维护套接字池的session错误:",err.Error())
		}
		mongoSession.SetMode(mgo.Monotonic,true)
		//使用指定数据库
		Db = mongoSession.DB(mongoConfig.GETDb())
	}

}
func GetMgo() *mgo.Session {
	return mongoSession
}

func GetDataBase() *mgo.Database {
	return Db
}
func GetErrNotFound() error {
	return mgo.ErrNotFound
}


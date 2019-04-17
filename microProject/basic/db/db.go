package db

import (

	"fmt"
	"microProject/basic/config"
	"github.com/micro/go-log"
	"github.com/jinzhu/gorm"

	"sync"
)

var (
	inited  bool
	mysqlDB *gorm.DB
	m       sync.RWMutex
)

// Init 初始化数据库
func Init() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[Init] db 已经初始化过")
		log.Logf(err.Error())
		return
	}

	// 如果配置声明使用mysql
	if config.GetMysqlConfig().GetEnabled() {
		initMysql()
	}

	inited = true
}

// GetDB 获取db
func GetDB() *gorm.DB{
	return mysqlDB
}

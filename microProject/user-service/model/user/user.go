package user

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
	"time"
	"microProject/basic/db"
	proto "microProject/user-service/proto/service"
)

var (
	s *service
	m sync.RWMutex
)

type User struct {
	ID        uint64      `gorm:"pk autoincr" json:"-"`
	Name        string    `gorm:"column:name" json:"user_name"`
	Pwd         string    `gorm:"column:pwd" json:"pwd"`
	CreatedTime time.Time `gorm:"Created"`
	UpdatedTime time.Time `gorm:"Updated"`
}

func QueryUserByNname(userName string) ([]*proto.User,  error) {
	list := make([]*proto.User,0)
	o := db.GetDB()
	err := o.Where("name=?", userName).Select("id,name,pwd").First(&list).Error
	/*result := map[string]interface{}{
		"id":list.ID,
		"user_name":list.Name,
		"pwd":list.Pwd,id
		"create_time":list.CreatedTime.Format("2006-01-02 15:04:05"),
	}*/
	if err != nil{
		fmt.Println("错误",err)
		return list,errors.Wrap(err,"错误")
	}

	return list,nil
}

// Service 用户服务类
/*type Service interface {
	// QueryUserByName 根据用户名获取用户
	QueryUserByName(userName string) (ret *proto.User, err error)
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}
*/
// service 服务
type service struct {

}


// Init 初始化用户服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	s = &service{}
}
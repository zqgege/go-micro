package phoneyzm

import (
	"fmt"
	"microProject/basic/mongo"
	"sync"
	"gopkg.in/mgo.v2/bson"
	"time"

	proto "microProject/user-service/proto/phoneyzm"
)


var (
	yzmService *PhoneYzmService
	m sync.RWMutex
)

type PhoneYzmService struct {

}

/*type PhoneYzm struct {
	Id bson.ObjectId `bson:"_id"`
	Phone  		int64      `bson:"phone"`
	Yzm    		int64      `bson:"yzm"`
	IsUse  		int64      `bson:"is_use"`
	ExTime      int64	   `bson:"ex_time"`
	CreateTime  int64	   `bson:"create_time"`
	MechanismId int64      `bson:"mechanism_id"`
}*/

func Init()  {
	m.Lock()
	defer m.Unlock()
	if yzmService != nil {
		return
	}
	yzmService = &PhoneYzmService{}

}
//获取根据手机号获取数据
func GetPhoneYzm(phone int64)(proto.PhoneYzm,error)  {
	//phoneyzm := make([]*proto.PhoneYzm,0)
	var phoneyzm proto.PhoneYzm
	fmt.Println(phone)
	con := mongo.GetDataBase().C("phoneyzm")
	if err :=con.Find(bson.M{"phone":phone}).One(&phoneyzm); err!=nil{
		if err.Error() != mongo.GetErrNotFound().Error(){
			return phoneyzm,err
		}
	}
	return phoneyzm,nil
}

//插入
func CreatePhoneYzm(phone int64,mechanismId int64)(_,err error)  {
	con := mongo.GetDataBase().C("phoneyzm")
	data := &proto.PhoneYzm{
		//Id:      bson.NewObjectId(),
		MechanismId:   mechanismId,
		Yzm:123456,
		CreateTime: time.Now().Unix(),
		Phone:phone,
	}
	err = con.Insert(data)
	if err != nil {
		//return fmt.Errorf(err.Error())
		fmt.Println("插入失败")
		panic(err)
	}
	fmt.Println("插入成功",err)
	return
}
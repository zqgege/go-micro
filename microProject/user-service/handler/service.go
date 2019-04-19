package handler

import (
	"bytes"
	"context"
	_ "encoding/json"
	"fmt"
	"log"
	py "microProject/user-service/model/phoneyzm"
	us "microProject/user-service/model/user"
	yzm "microProject/user-service/proto/phoneyzm"
	s "microProject/user-service/proto/service"
	"time"
)


type Service struct {

}

func Init()  {
	var err error
	//userService,err = us.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

func (e *Service) QueryUserByNname(ctx context.Context,req *s.Request,rep *s.Response) error  {
	list,err := us.QueryUserByNname(req.UserName)
	if err != nil {
		rep.Error = &s.Error{
			Code:500,
			Detail:err.Error(),
		}
		return err
	}
	fmt.Println("返回",list)
	rep.User = list

	return nil

}

func (e *Service)CreateUser(ctx context.Context,req *s.Request,rep *s.Response) error {
	user :=&us.User{
		Name:req.UserName,
		Pwd:req.UserPwd,
		CreatedTime:time.Now(),
		UpdatedTime:time.Now(),
	}
	err := us.CreateUser(user)
	if err != nil {
		return err
	}
	/*data := map[string]interface{}{
		"id":user.ID,
		"name":user.Name,
		"pwd":"",
	}*/
	data,err := us.GetById(user.ID)
	rep.User = data
	return nil
}

//根据手机号获取验证码
func (e *Service) GetPhoneYzm(ctx context.Context,req *yzm.Request,rep *yzm.Response)error  {
	list,err := py.GetPhoneYzm(req.Phone)
	if err != nil {
		rep.Error = &yzm.Error{
			Code:500,
			Datail:err.Error(),
		}
		return err
	}
	fmt.Println(list)
	rep.Yzm  = list.Yzm
	return nil
}

//创建验证码
func (e *Service)CreatePhoneYzm(ctx context.Context,req *yzm.Request,rep *yzm.Response) error  {
	res_yzm,err := py.CreatePhoneYzm(req.Phone,req.MechanismId)
	if err != nil {
		rep.Error = &yzm.Error{
			Code:500,
			Datail:err.Error(),
		}
		return err
	}
	fmt.Println(res_yzm)
	return nil
}
package handler

import (
	"context"
	_"encoding/json"
	"fmt"
	"log"
	s "microProject/user-service/proto/service"
	us "microProject/user-service/model/user"
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
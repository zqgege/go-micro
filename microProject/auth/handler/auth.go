package handler

import (
	"context"
	"microProject/auth/model/access"
	"strconv"

	"github.com/micro/go-log"

	 auth "microProject/auth/proto/auth"
)

var (
	accessService access.Service
)

func Init()  {
	var err error
	accessService,err = access.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误，%s", err)
		return
	}
}
type Service struct{}

// MakeAccessToken 生成token
func (s *Service) MakeAccessToken(ctx context.Context,req *auth.Request,rep *auth.Response) error {
	log.Log("[MakeAccessToken] 收到创建token请求")
	token,err := accessService.MakeAccessToken(&access.Subject{
		ID:strconv.FormatUint(req.UserId, 10),
		Name:req.UserName,
	})
	if err != nil {
		rep.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Logf("[MakeAccessToken], token生成失败，err：%s", err)
		return err
	}
	rep.Token = token
	return nil
}

// DelUserAccessToken 清除用户token
func (s *Service) DelUserAccessToken(ctx context.Context,req *auth.Request,rep *auth.Response)error {
	log.Log("[DelUserAccessToken] 清除用户token")
	err := accessService.DelUserAccessToken(req.Token)
	if err != nil {
		rep.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Logf("[DelUserAccessToken] 清除用户token失败，err：%s", err)
		return err
	}

	return nil
}

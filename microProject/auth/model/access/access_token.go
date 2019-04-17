package access

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-log/log"
	"github.com/micro/go-micro/broker"
	"microProject/basic/config"
	"time"
)

//生成与获取token的主要代码

var (
	// tokenExpiredDate app token过期日期 30天
	tokenExpiredDate = 3600 * 24 * 30 * time.Second
	// tokenIDKeyPrefix tokenID 前缀
	tokenIDKeyPrefix = "token:auth:id:"

	tokenExpiredTopic = "mu.micro.book.topic.auth.tokenExpired"

)

// Subject token 持有者
type Subject struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

// standardClaims token 标准的Claims
type standardClaims struct {
	SubjectID string `json:"subjectId,omitempty"`
	Name      string `json:"name,omitempty"`
	jwt.StandardClaims
}

//生成token并保存到redis中

func (s *service) MakeAccessToken(subject *Subject) (ret string,err error){
	m,err := s.createTokenClaims(subject)
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 创建token Claim 失败，err: %s", err)
	}

	//创建
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,m)
	ret,err = token.SignedString([]byte(config.GetJwtConfig().GetSecretKey()))
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 创建token失败，err: %s", err)
	}

	//保存到redis
	err = s.saveTokenToCachae(subject,ret)
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 保存token到缓存失败，err: %s", err)
	}
	return
}
// GetCachedAccessToken 获取token
func (s *service) GetCachedAccessToken(subject *Subject)(token string,err error) {
	token,err = s.getTokenFromCache(subject)
	if err != nil {
		return "", fmt.Errorf("[GetCachedAccessToken] 从缓存获取token失败，err: %s", err)
	}

	return
}
// DelUserAccessToken 清除用户token
func (s *service) DelUserAccessToken(tk string)(err error) {
	// 解析token字符串
	claims, err := s.parseToken(tk)
	if err != nil {
		return fmt.Errorf("[DelUserAccessToken] 错误的token，err: %s", err)
	}
	// 通过解析到的用户id删除
	err = s.delTokenFromCache(&Subject{
		ID: claims.Subject,
	})

	if err != nil {
		return fmt.Errorf("[DelUserAccessToken] 清除用户token，err: %s", err)
	}

	// 广播删除
	msg := &broker.Message{
		Body: []byte(claims.Subject),
	}
	if err := broker.Publish(tokenExpiredTopic, msg); err != nil {
		log.Logf("[pub] 发布token删除消息失败： %v", err)
	} else {
		fmt.Println("[pub] 发布token删除消息：", string(msg.Body))
	}

	return
}

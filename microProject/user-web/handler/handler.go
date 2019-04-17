package handler

import (
	"context"
	"encoding/json"
	"github.com/go-log/log"
	"github.com/micro/go-micro/client"
	auto "microProject/auth/proto/auth"
	us "microProject/user-service/proto/service"
	"net/http"
	"time"
)

var (
	serviceClient us.UserServicesService
	autoCli auto.Service
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init()  {
	serviceClient = us.NewUserServicesService("mu.micro.book.srv.user",client.DefaultClient)
	autoCli = auto.NewService("mu.micro.book.srv.auth",client.DefaultClient)
}

func Login(w http.ResponseWriter,r *http.Request)  {
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w,"非法请求",400)
		return
	}

	r.ParseForm()

	//调用后台服务

	rsp,err := serviceClient.QueryUserByNname(context.TODO(),&us.Request{
		UserName:r.Form.Get("userName"),
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := map[string]interface{}{
		"ref" : time.Now().UnixNano(),
	}
	//var user  map[string]string
	var id uint64
	var name  string
	var pwd string

	for _, u := range rsp.User {
		id = uint64(u.Id)
		name = u.Name
		pwd = u.Pwd
		u.Pwd = ""
		/*user["id"] = string(u.Id)
		user["name"] = u.Name
		user["pwd"] = u.Pwd*/
	}

	if pwd == r.Form.Get("pwd"){
		response["success"] = rsp.Success
		response["data"] = rsp.User
		log.Logf("[Login] 密码校验完成，生成token...")
		//生成token
		rsp2,err :=autoCli.MakeAccessToken(context.TODO(),&auto.Request{
			UserId:id,
			UserName:name,
		})
		if err != nil {
			log.Logf("[Login] 创建token失败，err：%s", err)
			http.Error(w, err.Error(), 500)
			return
		}
		log.Logf("[Login] token %s", rsp2.Token)
		response["token"] = rsp2.Token
		// 同时将token写到cookies中
		w.Header().Add("set-cookie", "application/json; charset=utf-8")
		// 过期30分钟
		expire := time.Now().Add(30 * time.Minute)
		cookie := http.Cookie{Name: "remember-me-token", Value: rsp2.Token, Path: "/", Expires: expire, MaxAge: 90000}
		http.SetCookie(w, &cookie)

	} else {
		response["success"] = false
		response["error"] = &Error{
			Detail:"密码错误",
		}
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w,err.Error(),500)
		return
	}
}

func Logout(w http.ResponseWriter,r *http.Request)  {
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w,"非法请求",400)
		return
	}

	tokenCookie,err := r.Cookie("remember-me-token")
	if err != nil {
		log.Logf("token获取失败")
		http.Error(w, "非法请求", 400)
		return
	}
	//删除token
	_, err = autoCli.DelUserAccessToken(context.TODO(), &auto.Request{
		Token:tokenCookie.Value,
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//清楚cookie
	cookie := http.Cookie{
		Name:"remember-me-token",
		Value:"",
		Path:"/",
		Expires:time.Now().Add(0*time.Second),
		MaxAge:0,
	}
	http.SetCookie(w,&cookie)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回结果
	response := map[string]interface{}{
		"ref":     time.Now().UnixNano(),
		"success": true,
	}

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
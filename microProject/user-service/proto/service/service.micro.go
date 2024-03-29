// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/service/service.proto

package mu_micro_book_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserServices service

type UserServicesService interface {
	QueryUserByNname(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	CreateUser(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type userServicesService struct {
	c    client.Client
	name string
}

func NewUserServicesService(name string, c client.Client) UserServicesService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "mu.micro.book.srv.user"
	}
	return &userServicesService{
		c:    c,
		name: name,
	}
}

func (c *userServicesService) QueryUserByNname(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserServices.QueryUserByNname", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesService) CreateUser(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserServices.CreateUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserServices service

type UserServicesHandler interface {
	QueryUserByNname(context.Context, *Request, *Response) error
	CreateUser(context.Context, *Request, *Response) error
}

func RegisterUserServicesHandler(s server.Server, hdlr UserServicesHandler, opts ...server.HandlerOption) error {
	type userServices interface {
		QueryUserByNname(ctx context.Context, in *Request, out *Response) error
		CreateUser(ctx context.Context, in *Request, out *Response) error
	}
	type UserServices struct {
		userServices
	}
	h := &userServicesHandler{hdlr}
	return s.Handle(s.NewHandler(&UserServices{h}, opts...))
}

type userServicesHandler struct {
	UserServicesHandler
}

func (h *userServicesHandler) QueryUserByNname(ctx context.Context, in *Request, out *Response) error {
	return h.UserServicesHandler.QueryUserByNname(ctx, in, out)
}

func (h *userServicesHandler) CreateUser(ctx context.Context, in *Request, out *Response) error {
	return h.UserServicesHandler.CreateUser(ctx, in, out)
}

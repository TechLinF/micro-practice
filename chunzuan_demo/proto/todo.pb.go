// Code generated by protoc-gen-go.
// source: todo.proto
// DO NOT EDIT!

/*
Package service_todo is a generated protocol buffer package.

It is generated from these files:
	todo.proto

It has these top-level messages:
	AddTaskReq
	Null
*/
package service_todo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddTaskReq struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	OwnerId     string `protobuf:"bytes,2,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
}

func (m *AddTaskReq) Reset()                    { *m = AddTaskReq{} }
func (m *AddTaskReq) String() string            { return proto.CompactTextString(m) }
func (*AddTaskReq) ProtoMessage()               {}
func (*AddTaskReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Null struct {
}

func (m *Null) Reset()                    { *m = Null{} }
func (m *Null) String() string            { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()               {}
func (*Null) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*AddTaskReq)(nil), "service.todo.AddTaskReq")
	proto.RegisterType((*Null)(nil), "service.todo.Null")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Todo service

type TodoClient interface {
	AddTask(ctx context.Context, in *AddTaskReq, opts ...client.CallOption) (*Null, error)
}

type todoClient struct {
	c           client.Client
	serviceName string
}

func NewTodoClient(serviceName string, c client.Client) TodoClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "service.todo"
	}
	return &todoClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *todoClient) AddTask(ctx context.Context, in *AddTaskReq, opts ...client.CallOption) (*Null, error) {
	req := c.c.NewRequest(c.serviceName, "Todo.AddTask", in)
	out := new(Null)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Todo service

type TodoHandler interface {
	AddTask(context.Context, *AddTaskReq, *Null) error
}

func RegisterTodoHandler(s server.Server, hdlr TodoHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Todo{hdlr}, opts...))
}

type Todo struct {
	TodoHandler
}

func (h *Todo) AddTask(ctx context.Context, in *AddTaskReq, out *Null) error {
	return h.TodoHandler.AddTask(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc9, 0x4f, 0xc9,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x29, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5,
	0x03, 0x89, 0x29, 0xe5, 0x73, 0x71, 0x39, 0xa6, 0xa4, 0x84, 0x24, 0x16, 0x67, 0x07, 0xa5, 0x16,
	0x0a, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x01, 0x59, 0x42,
	0x92, 0x5c, 0x1c, 0xf9, 0xe5, 0x79, 0xa9, 0x45, 0xf1, 0x40, 0x51, 0x26, 0xb0, 0x28, 0x3b, 0x98,
	0xef, 0x99, 0x22, 0x24, 0xc2, 0xc5, 0x5a, 0x92, 0x59, 0x92, 0x93, 0x2a, 0xc1, 0x0c, 0x16, 0x87,
	0x70, 0x84, 0x14, 0xb8, 0xb8, 0x53, 0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0x32, 0xf3, 0xf3,
	0x24, 0x58, 0xc0, 0x72, 0xc8, 0x42, 0x4a, 0x6c, 0x5c, 0x2c, 0x7e, 0xa5, 0x39, 0x39, 0x46, 0x8e,
	0x5c, 0x2c, 0x21, 0x40, 0x07, 0x08, 0x59, 0x72, 0xb1, 0x43, 0x1d, 0x20, 0x24, 0xa1, 0x87, 0xec,
	0x34, 0x3d, 0x84, 0xbb, 0xa4, 0x84, 0x50, 0x65, 0x40, 0x06, 0x28, 0x31, 0x24, 0xb1, 0x81, 0x3d,
	0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x68, 0xb8, 0x7b, 0xca, 0xde, 0x00, 0x00, 0x00,
}
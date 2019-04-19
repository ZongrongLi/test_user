/*
 * File: main.go
 * Project: go-rpc
 * File Created: Friday, 5th April 2019 12:00:35 am
 * Author: lizongrong (389006500@qq.com)
 * -----
 * Last Modified: Friday, 5th April 2019 4:48:07 pm
 * Modified By: lizongrong (389006500@qq.com>)
 * -----
 * null lizongrong - 2019
 */
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/docker/libkv/store"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/tiancai110a/go-rpc/registry"
	"github.com/tiancai110a/go-rpc/registry/libkv"

	"github.com/golang/glog"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/tiancai110a/go-rpc/protocol"
	"github.com/tiancai110a/go-rpc/server"
	"github.com/tiancai110a/go-rpc/service"
	"github.com/tiancai110a/go-rpc/transport"
)

func testMiddleware1(rw *http.ResponseWriter, r *http.Request, c *server.Middleware) {

	fmt.Println("before===testMiddlewarec1")
	c.Next(nil, nil)

	fmt.Println("after===testMiddlewarec1")
}

func testMiddleware2(rw *http.ResponseWriter, r *http.Request, c *server.Middleware) {
	fmt.Println("before===testMiddlewarec2")
	c.Next(nil, nil)

	fmt.Println("after===testMiddlewarec2")
}

func testMiddleware3(rw *http.ResponseWriter, r *http.Request, c *server.Middleware) {
	fmt.Println("before===testMiddlewarec3")
	c.Next(nil, nil)
	fmt.Println("after===testMiddlewarec3")
}
func TestAdd(ctx context.Context, resp *service.Resp) {

	glog.Info("===========================================================================================resultful func")
	glog.Info("==================================test1:", ctx.Value("test1"))
	glog.Info("==================================test:", ctx.Value("test"))
	glog.Info("==================================name:", ctx.Value("name"))
	glog.Info("==================================pass:", ctx.Value("pass"))
	//	res.data

	resp.Add("name", "tiancai")
	resp.Add("res1", "3.14")
	resp.Add("list1", "1234,4567,1234,0987,3333")
	return
}

//用来停止server，测试心跳功能
var gs server.RPCServer

func StartServer(op *server.Option) {
	go func() {
		s, err := server.NewSGServer(op)
		if err != nil {
			glog.Error("new serializer failed", err)
			return
		}
		//s.Register(service.TestService{})
		err = s.Register(service.ArithService{})

		gs = s
		if err != nil {
			glog.Error("Register failed,err:", err)

		}

		sk := s.Group(service.POST, "/v1/invoke/")
		if sk == nil {
			glog.Error("server dose not implement http server")
			return
		}
		sk.Route("/Add", TestAdd)
		s.Use(testMiddleware1)
		s.Use(testMiddleware2)
		s.Use(testMiddleware3)
		go s.Serve("tcp", "127.0.0.1:8889", nil)
	}()
}

func main() {

	opentracing.SetGlobalTracer(mocktracer.New())

	//单机伪集群
	r1 := libkv.NewKVRegistry(store.ZK, "my-app", "/root/lizongrong/service",
		[]string{"127.0.0.1:1181", "127.0.0.1:2181", "127.0.0.1:3181"}, 1e10, nil)
	servertOption := server.Option{
		ProtocolType:   protocol.Default,
		SerializeType:  protocol.SerializeTypeMsgpack,
		CompressType:   protocol.CompressTypeNone,
		TransportType:  transport.TCPTransport,
		ShutDownWait:   time.Second * 12,
		Registry:       r1,
		RegisterOption: registry.RegisterOption{"my-app"},
		Tags:           map[string]string{"idc": "lf"}, //只允许机房为lf的请求，客户端取到信息会自己进行转移
		HttpServePort:  5080,
	}

	StartServer(&servertOption)
	time.Sleep(time.Second * 265)

}

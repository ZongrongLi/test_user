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
	"strconv"
	"strings"
	"time"

	"github.com/docker/libkv/store"
	"github.com/lexkong/log"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/tiancai110a/go-rpc/registry"
	"github.com/tiancai110a/go-rpc/registry/libkv"
	"github.com/tiancai110a/test_user/config"

	"github.com/golang/glog"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/tiancai110a/go-rpc/protocol"
	"github.com/tiancai110a/go-rpc/server"
	"github.com/tiancai110a/go-rpc/transport"
)

//用来停止server，测试心跳功能
var gs server.RPCServer

func StartServer(op *server.Option) {
	go func() {
		s, err := server.NewSGServer(op)
		if err != nil {
			glog.Error("new serializer failed", err)
			return
		}

		Load(s)
		go s.Serve("tcp", viper.GetString("tcpurl"), nil)
	}()
}

func main() {

	opentracing.SetGlobalTracer(mocktracer.New())

	if err := config.Init(""); err != nil {
		panic(err)
	}

	var r1 registry.Registry
	if viper.GetString("discovery.name") == "zk" {
		nodes := viper.GetString("discovery.nodes")
		zknode := strings.Split(nodes, ",")
		interval, err := strconv.ParseFloat(viper.GetString("discovery.updateinterval"), 64)
		if err != nil {
			log.Infof("parse interval err: %s", err)
			interval = 1e10
		}

		r1 = libkv.NewKVRegistry(store.ZK, viper.GetString("server_name"), viper.GetString("discovery.path"),
			zknode, time.Duration(interval), nil)

	} else {
		glog.Error("discovery is not set")
		return
	}
	servertOption := server.Option{
		ProtocolType:   protocol.Default,
		SerializeType:  protocol.SerializeTypeMsgpack,
		CompressType:   protocol.CompressTypeNone,
		TransportType:  transport.TCPTransport,
		ShutDownWait:   time.Second * 12,
		Registry:       r1,
		RegisterOption: registry.RegisterOption{viper.GetString("server_name")},
		Tags:           map[string]string{"idc": viper.GetString("idc")}, //只允许机房为lf的请求，客户端取到信息会自己进行转移
		HttpServeOpen:  false,
	}

	StartServer(&servertOption)
	time.Sleep(time.Second * 265)

}

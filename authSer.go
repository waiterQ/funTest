package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/waiterQ/funTest/gen-go/auth"
	"os"
)

const (
	NetworkAddr = "127.0.0.1:9005"
)

type AuthSer struct{}

func (*AuthSer) Register(name string, pwd string) (res string, err error) {
	res = name + "+" + pwd
	fmt.Println(res)
	return
}

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	handler := &AuthSer{}
	processor := auth.NewAuthSerProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("authThrift server in :" + NetworkAddr)
	server.Serve()
}

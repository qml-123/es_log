package main

import (
	"fmt"

	"github.com/qml-123/es_log/gen-go/es_log"
	"github.com/qml-123/es_log/handler"

	"github.com/apache/thrift/lib/go/thrift"
)

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	addr := "localhost:9090"
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		fmt.Println("Error creating server socket:", err)
		return
	}

	server := thrift.NewTSimpleServer4(
		es_log.NewLogServiceProcessor(&handler.LogHandler{}),
		transport,
		transportFactory,
		protocolFactory,
	)
	fmt.Println("Starting RPC server on", addr)
	if err := server.Serve(); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}

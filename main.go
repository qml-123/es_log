package main

import (
	"fmt"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
	"github.com/qml-123/GateWay/common"
	es_log "github.com/qml-123/es_log/kitex_gen/es_log/logservice"
)

const (
	configPath = "config/services.json"
)

func main() {
	conf, err := common.GetJsonFromFile(configPath)
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:"+fmt.Sprintf("%d", conf.ListenPort))
	if err != nil {
		panic(err)
	}
	svr := es_log.NewServer(new(LogServiceImpl), server.WithServiceAddr(addr))

	addr, _ = net.ResolveTCPAddr("tcp", conf.ListenIp+":"+fmt.Sprintf("%d", conf.ListenPort))
	if err = common.InitConsul(addr, conf); err != nil {
		panic(err)
	}

	defer common.CloseConsul(addr, conf)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}

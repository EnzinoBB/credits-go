package core

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/EnzinoBB/credits-go/api"
	"github.com/EnzinoBB/credits-go/general"
	"github.com/EnzinoBB/credits-go/model"

	"github.com/akamensky/base58"
	"github.com/apache/thrift/lib/go/thrift"
)

type NodeClient struct {
	trans  thrift.TTransport
	client *api.APIClient
}

func NewNodeClient(nodeAddr string) (*NodeClient, error) {

	framed := false
	buffered := true
	addr := nodeAddr
	secure := false

	cfg := &thrift.TConfiguration{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		ConnectTimeout: time.Second * 10,
		SocketTimeout:  time.Second * 10,
	}

	var protocolFactory thrift.TProtocolFactory = thrift.NewTBinaryProtocolFactoryConf(cfg)

	var transportFactory thrift.TTransportFactory

	if buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	if framed {
		transportFactory = thrift.NewTFramedTransportFactoryConf(transportFactory, cfg)
	}

	var err error
	var transport thrift.TTransport
	if secure {
		transport, err = thrift.NewTSSLSocketConf(addr, cfg)
	} else {
		transport, err = thrift.NewTSocketConf(addr, cfg)
	}
	if err != nil {
		panic(err)
	}

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		panic(err)
	}

	if err := transport.Open(); err != nil {
		panic(err)
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	client := api.NewAPIClient(thrift.NewTStandardClient(iprot, oprot))

	node := &NodeClient{
		client: client,
		trans:  transport,
	}

	return node, nil
}

func (client *NodeClient) CloseConnection() {

	if client.trans.IsOpen() {
		client.trans.Close()
	}

}

func (client *NodeClient) CheckConnection() bool {

	return client.trans.IsOpen()

}

func (client *NodeClient) GetWalletData(wallet string) (*model.WalletData, error) {

	argvalue0, err := base58.Decode(wallet)
	if err != nil {
		return nil, err
	}
	value0 := general.Address(argvalue0)

	result, err := client.client.WalletDataGet(context.Background(), value0)
	if err != nil {
		return nil, err
	}

	return model.GetWalletData_Out(result.WalletData), nil
}

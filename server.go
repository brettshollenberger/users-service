package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	// mthrift "github.com/brettshollenberger/minion/thrift"
	"github.com/brettshollenberger/thrift-example/gen-go/users"
	. "github.com/brettshollenberger/users-service/handlers"
)

const (
	USERS_SERVICE string = "users"
)

func runServer(addr string) error {
	var transport thrift.TServerTransport
	var err error

	transport, err = thrift.NewTServerSocket(addr)

	if err != nil {
		return err
	}

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	fmt.Printf("%T\n", transport)

	usersHandler := NewUsersHandler()
	// processor := mthrift.NewTMultiplexedProcessor()
	// listener := mthrift.NewListener()

	// processor.RegisterProcessor(USERS_SERVICE, *users.NewUsersServiceProcessor(usersHandler))
	processor := users.NewUsersServiceProcessor(usersHandler)

	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", addr)

	return server.Serve()
}

func main() {
	runServer("127.0.0.1:9090")
}

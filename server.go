package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/brettshollenberger/thrift-example/gen-go/users"
	. "github.com/brettshollenberger/users-service/handlers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func runServer(addr string) error {
	var transport thrift.TServerTransport
	var err error
	var db gorm.DB

	transport, err = thrift.NewTServerSocket(addr)

	if err != nil {
		return err
	}

	db, err = gorm.Open("mysql", "root:@/users_service?charset=utf8&parseTime=True&loc=Local")

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	usersHandler := NewUsersHandler(db)
	processor := users.NewUsersServiceProcessor(usersHandler)

	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", addr)

	return server.Serve()
}

func main() {
	runServer("127.0.0.1:9090")
}

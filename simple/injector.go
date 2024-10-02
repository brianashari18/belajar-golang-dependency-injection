//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository,
		NewSimpleService,
	)

	return nil, nil
}

func InitializeDatabase() *DatabaseRepository {
	wire.Build(NewDatabaseMongoDB, NewDatabasePostgreSQL, NewDatabaseRepository)

	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)

	return nil
}

// Untuk versi return interfacenya
func InitializeHelloService() *HelloService {
	wire.Build(NewHelloService, NewSayHelloImpl)

	return nil
}

// Untuk versi return impl nya pake interface binding
//var helloSet = wire.NewSet(
//	NewSayHelloImpl,
//	wire.Bind(new(SayHello), new(*SayHelloImpl)),
//)
//
//func InitializeHelloService() *HelloService {
//	wire.Build(NewHelloService, helloSet)
//	return nil
//}

var fooValue = &Foo{}
var barValue = &Bar{}

func InitializeFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))

	return nil
}

func InitializeReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

// Struct Field Provider
func InitializeConfiguration() *Configuration {
	wire.Build(NewApplication, wire.FieldsOf(new(*Application), "Configuration"))

	return nil
}

// Cleanup Function
func InitializeConnection(name string) (*Connection, func()) {
	wire.Build(NewFile, NewConnection)
	return nil, nil
}

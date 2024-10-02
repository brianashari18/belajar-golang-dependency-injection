// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

// Injectors from injector.go:

func InitializedService(isError bool) (*SimpleService, error) {
	simpleRepository := NewSimpleRepository(isError)
	simpleService, err := NewSimpleService(simpleRepository)
	if err != nil {
		return nil, err
	}
	return simpleService, nil
}

func InitializeDatabase() *DatabaseRepository {
	databasePostgreSQL := NewDatabasePostgreSQL()
	databaseMongoDB := NewDatabaseMongoDB()
	databaseRepository := NewDatabaseRepository(databasePostgreSQL, databaseMongoDB)
	return databaseRepository
}

func InitializeFooBarService() *FooBarService {
	fooRepository := NewFooRepository()
	fooService := NewFooService(fooRepository)
	barRepository := NewBarRepository()
	barService := NewBarService(barRepository)
	fooBarService := NewFooBarService(fooService, barService)
	return fooBarService
}

// Untuk versi return interfacenya
func InitializeHelloService() *HelloService {
	sayHello := NewSayHelloImpl()
	helloService := NewHelloService(sayHello)
	return helloService
}

func InitializeFooBarUsingValue() *FooBar {
	foo := _wireFooValue
	bar := _wireBarValue
	fooBar := &FooBar{
		Foo: foo,
		Bar: bar,
	}
	return fooBar
}

var (
	_wireFooValue = fooValue
	_wireBarValue = barValue
)

func InitializeReader() io.Reader {
	reader := _wireFileValue
	return reader
}

var (
	_wireFileValue = os.Stdin
)

// Struct Field Provider
func InitializeConfiguration() *Configuration {
	application := NewApplication()
	configuration := application.Configuration
	return configuration
}

// Cleanup Function
func InitializeConnection(name string) (*Connection, func()) {
	file, cleanup := NewFile(name)
	connection, cleanup2 := NewConnection(file)
	return connection, func() {
		cleanup2()
		cleanup()
	}
}

// injector.go:

var fooSet = wire.NewSet(NewFooRepository, NewFooService)

var barSet = wire.NewSet(NewBarRepository, NewBarService)

var fooValue = &Foo{}

var barValue = &Bar{}

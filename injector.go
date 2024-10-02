//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"golang-dependency-injection/app"
	"golang-dependency-injection/controller"
	"golang-dependency-injection/middleware"
	"golang-dependency-injection/repository"
	"golang-dependency-injection/service"
	"net/http"
)

func InitializeServer() *http.Server {
	wire.Build(
		app.NewDB,
		NewValidator,
		repository.NewCategoryRepository,
		service.NewCategoryService,
		controller.NewCategoryController,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}

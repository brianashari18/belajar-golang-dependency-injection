package main

import (
	"golang-dependency-injection/helper"
	"golang-dependency-injection/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	return validator.New() // Add options here if necessary
}

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializeServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

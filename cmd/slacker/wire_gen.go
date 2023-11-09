// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"slacker/internal/biz/record"
	"slacker/internal/biz/user"
	"slacker/internal/conf"
	"slacker/internal/data"
	"slacker/internal/server"
	"slacker/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(logger log.Logger, confServer *conf.Server, confData *conf.Data, auth *conf.Auth) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	repo := data.NewUserRepo(dataData, logger)
	useCase := user.NewUseCase(logger, repo)
	userService := service.NewUserService(logger, useCase, auth)
	recordRepo := data.NewRecordRepo(dataData, logger)
	recordUseCase := record.NewUseCase(logger, recordRepo)
	recordService := service.NewRecordService(logger, recordUseCase)
	httpServer := server.NewHTTPServer(logger, confServer, auth, userService, recordService)
	app := newApp(logger, httpServer)
	return app, func() {
		cleanup()
	}, nil
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"staff_client/internal/biz"
	"staff_client/internal/conf"
	"staff_client/internal/data"
	"staff_client/internal/server"
	"staff_client/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, auth *conf.Auth, confService *conf.Service, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	staffClient := data.NewStaffServiceClient(auth, confService, discovery)
	client := data.NewRedis(confData)
	dataData, err := data.NewData(confData, logger, staffClient, client)
	if err != nil {
		return nil, nil, err
	}
	staffClientRepo := data.NewStaffClientRepo(dataData, logger)
	staffClientUsecase := biz.NewStaffClientUsecase(staffClientRepo, logger, auth)
	staffClientService := service.NewStaffClientService(staffClientUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, auth, staffClientService, logger)
	registrar := data.NewRegistrar(registry)
	app := newApp(logger, httpServer, registrar)
	return app, func() {
	}, nil
}

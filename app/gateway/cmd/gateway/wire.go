//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos_sample/app/device/internal/server"
	"kratos_sample/app/gateway/internal/router"
	"kratos_sample/app/gateway/internal/server"
)

func wireApp(
	*conf.Server,
	*conf.Etcd,
	*conf.Service,
	*conf.Data,
	log.Logger,
) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		//service.ProviderSet,
		router.ProviderSet,
		newApp,
	))
}

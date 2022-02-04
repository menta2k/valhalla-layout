//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/menta2l/valhalla-layout/internal/biz"
	"github.com/menta2l/valhalla-layout/internal/conf"
	"github.com/menta2l/valhalla-layout/internal/data"
	"github.com/menta2l/valhalla-layout/internal/server"
	"github.com/menta2l/valhalla-layout/internal/service"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Auth, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}

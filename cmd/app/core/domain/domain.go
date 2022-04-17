package domain

import (
	"context"
	"github.com/Feoks/loader/cmd/app/core/infrastructure"
	"net/http"

	"github.com/Feoks/loader/pkg/echo"
	"github.com/Feoks/loader/pkg/health"
	googlegrpc "google.golang.org/grpc"
)

var serviceLayer = serviceNetwork{}

func init() {
	serviceLayer = serviceNetwork{
		httpHandler: make(map[string]http.Handler, 0),
		grpcHandler: make([]func(server *googlegrpc.Server), 0, 0),
	}
}

//public
type serviceNetwork struct {
	httpHandler map[string]http.Handler
	grpcHandler []func(*googlegrpc.Server)
}

type EchoService interface {
	HttpHandlers() map[string]http.Handler
	GrpcHandlers() []func(*googlegrpc.Server)
	Close()
}

//private
type service struct {
	infra infrastructure.Infrastructure

	healthService health.Service
	echoService   echo.Service
}

type initService interface {
	healthService
	echoService

	initNetworkHandler(ctx context.Context)
	initDBHandler(ctx context.Context)
}

func NewDomain(ctx context.Context, infra infrastructure.Infrastructure) EchoService {

	go func(s service) initService {
		s.initHealthService(ctx)
		s.initEchoService(NewEchoRepository())
		s.initNetworkHandler(ctx)
		return &s
	}(service{infra: infra})

	return &serviceLayer
}

func (s *service) initDBHandler(ctx context.Context) {
	//TODO Доделать подключение к БД
	panic("implement me")
}

func (s *service) initNetworkHandler(ctx context.Context) {
	//make HTTP handler
	{
		serviceLayer.httpHandler[HEALTH_PATH] = health.MakeHTTPHandler(ctx, s.getHealth())
		serviceLayer.httpHandler[ECHO_PATH] = echo.MakeHTTPHandler(ctx, s.getEcho())
	}

	//make GRPC handler
	{
		serviceLayer.grpcHandler = append(
			serviceLayer.grpcHandler,
			health.JoinGRPC(ctx, s.getHealth()),
			echo.JoinGRPC(ctx, s.getEcho()),
		)
	}
}

func (s *serviceNetwork) HttpHandlers() map[string]http.Handler {
	return s.httpHandler
}

func (s *serviceNetwork) GrpcHandlers() []func(*googlegrpc.Server) {
	return s.grpcHandler
}

func (s *serviceNetwork) Close() {
	//TODO implement me
	panic("implement me")
}

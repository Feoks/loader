package domain

import (
	e "github.com/Feoks/loader/internal/repository/echo"
	"github.com/Feoks/loader/pkg/echo"
)

const ECHO_PATH = "/echo"

type echoService interface {
	initEchoService(repo e.Echo)
	getEcho() echo.Service
}

func (s *service) initEchoService(repo e.Echo) {
	echoService := echo.NewEchoService(repo)
	if s.infra.Configs().Sentry.Enabled {
		echoService = echo.NewSentryService(echoService)
	}
	s.echoService = echoService
}

func (s *service) getEcho() echo.Service {
	return s.echoService
}

package main

import (
	"git.repo.services.lenvendo.ru/grade-factor/echo/cmd/app/core"
	"git.repo.services.lenvendo.ru/grade-factor/echo/internal/server"
	"git.repo.services.lenvendo.ru/grade-factor/echo/tools/logging"
	"github.com/go-kit/kit/log/level"

	"os"
)

func ServerRun111(sc core.ServiceCore, ic core.InfrastructureCore) {

	ctx := ic.Ctx()
	logger := logging.FromContext(ctx)

	s, err := server.NewServer(
		server.SetConfig(ic.Cfg()),
		server.SetLogger(logger),
		server.SetHandler(sc.Service().HttpHandlers()),
		server.SetGRPC(sc.Service().GrpcHandlers()...),
	)
	if err != nil {
		level.Error(logger).Log("init", "server", "err", err)
		os.Exit(1)
	}
	defer s.Close()

	if err := s.AddHTTP(); err != nil {
		level.Error(logger).Log("err", err)
		os.Exit(1)
	}

	if err = s.AddGRPC(); err != nil {
		level.Error(logger).Log("err", err)
		os.Exit(1)
	}

	if err = s.AddMetrics(); err != nil {
		level.Error(logger).Log("err", err)
		os.Exit(1)
	}

	s.AddSignalHandler()
	s.Run()
}

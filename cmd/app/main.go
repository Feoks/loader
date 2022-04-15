package main

import (
	"context"
	c "git.repo.services.lenvendo.ru/grade-factor/echo/cmd/app/core"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serviceCore, infraCore := c.InitCore(ctx)
	defer infraCore.Closer()

	ServerRun111(serviceCore, infraCore)
}

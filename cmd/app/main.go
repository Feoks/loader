package main

import (
	"context"
	c "github.com/Feoks/loader/cmd/app/core"
	"github.com/Feoks/loader/internal/file_loader"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serviceCore, infraCore := c.InitCore(ctx)
	defer infraCore.Closer()

	file_loader.RunWatcher(infraCore)

	ServerRun111(serviceCore, infraCore)
}

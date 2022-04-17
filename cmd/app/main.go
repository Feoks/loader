package main

import (
	"context"
	c "github.com/Feoks/loader/cmd/app/core"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serviceCore, infraCore := c.InitCore(ctx)
	defer infraCore.Closer()

	ServerRun111(serviceCore, infraCore)
}

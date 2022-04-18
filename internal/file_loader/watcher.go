package file_loader

import (
	"fmt"
	"github.com/Feoks/loader/cmd/app/core"
	"github.com/Feoks/loader/tools/logging"
	"github.com/go-kit/kit/log/level"
	"os"
	"time"
)

// RunWatcher запускает в цикле r с паузами p секунд
func RunWatcher(ic core.InfrastructureCore) {
	cfg := ic.Cfg()
	ctx := ic.Ctx()
	logger := logging.FromContext(ctx)
	pause := cfg.Watcher.Period
	fileName := "upload/data.json" //cfg.Watcher.FileName
	loader := initLoader(fileName)
	fmt.Println(fileName)
	for {
		if err := loader.Load(fileName); err != nil {
			_ = level.Error(logger).Log("err", err)
			os.Exit(1)
		}
		time.Sleep(time.Duration(pause) * time.Second)
	}
}

package infrastructure

import (
	"context"
	"fmt"
	"os"

	"git.repo.services.lenvendo.ru/grade-factor/echo/configs"
	"git.repo.services.lenvendo.ru/grade-factor/echo/tools/logging"
	"github.com/go-kit/kit/log/level"
)

//Infrastructure - Context - возвращает актуальный context содержащий все инициализируемые инфраструктурные юниты
// Configs - конфигы из env, toml, bash или структур из /configs/config.go
// AllCloser - закрытие всех соединение, если закрытие предполагает инфраструктурным юнитом
type Infrastructure interface {
	Context() context.Context
	Configs() *configs.Config
	AllCloser()
}

//infrastructure ctx - контекст агрегирует в себе юниты
// актуальные конфиги
type infrastructure struct {
	ctx context.Context
	cfg *configs.Config
	bs  *bootstrap
}

//NewInfrastructure подключение всех юнитов
func NewInfrastructure(ctx context.Context) Infrastructure {

	cfg, err := InitConfig()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	b := &bootstrap{}
	for _, v := range b.initBootstrap(cfg.Configs()) {
		if ctx, err = v.init(ctx, cfg.Configs()); err != nil {
			level.Error(logging.FromContext(ctx)).Log("err", err, "msg", "")
		}
	}

	return &infrastructure{
		ctx: ctx,
		cfg: cfg.Configs(),
		bs:  b,
	}
}

func (infra *infrastructure) Context() context.Context {
	return infra.ctx
}

func (infra *infrastructure) Configs() *configs.Config {
	return infra.cfg
}

func (infra *infrastructure) AllCloser() {
	for _, v := range infra.bs.units {
		v.closer()
	}
}

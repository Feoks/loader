package infrastructure

import (
	"git.repo.services.lenvendo.ru/grade-factor/echo/configs"
	"github.com/pkg/errors"
)

type config struct {
	cfg *configs.Config
}

func InitConfig() (c *config, err error) {
	c = &config{cfg: configs.NewConfig()}
	if err := c.cfg.Read(); err != nil {
		return nil, errors.Wrap(err, "read config")
	}
	if err := c.cfg.Print(); err != nil {
		return nil, errors.Wrap(err, "read config")

	}
	return c, nil
}

func (c *config) Configs() *configs.Config {
	return c.cfg
}

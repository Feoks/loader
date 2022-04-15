package domain

import (
	e "git.repo.services.lenvendo.ru/grade-factor/echo/internal/repository/echo"
)

func NewEchoRepository() e.Echo {
	return e.NewEcho()
}

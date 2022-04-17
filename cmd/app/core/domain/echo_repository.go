package domain

import (
	e "github.com/Feoks/loader/internal/repository/echo"
)

func NewEchoRepository() e.Echo {
	return e.NewEcho()
}

package echo

import (
	"context"
	"git.repo.services.lenvendo.ru/grade-factor/echo/internal/repository/echo"
)

type echoService struct {
	e echo.Echo
}

func NewEchoService(e echo.Echo) Service {
	return &echoService{e}
}

func (s *echoService) GetEcho(ctx context.Context, req *GetEchoListRequest) (resp *GetEchoListResponse, err error) {
	a := GetEchoListResponse{
		Echos: []Echo{
			{
				Id:       uint32(1),
				Title:    "title",
				Reminder: s.e.Ping(),
			},
		},
	}
	if err != nil {
		return &a, err
	}
	return &a, nil
}

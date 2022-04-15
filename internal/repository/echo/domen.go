package echo

type echo struct{}

type Echo interface {
	Ping() string
}

func NewEcho() Echo {
	return &echo{}
}

func (e *echo) Ping() string {
	return "ok"
}

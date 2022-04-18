package file_loader

import (
	"github.com/Feoks/echo/pkg/task"
)

type Sender interface {
	Send(tasks *[]task.Task) error
}

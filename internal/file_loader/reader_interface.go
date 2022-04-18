package file_loader

import (
	"github.com/Feoks/echo/pkg/task"
)

type Reader interface {
	Read() (*[]task.Task, error)
}

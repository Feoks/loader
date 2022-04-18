package file_loader

import (
	"fmt"
	"github.com/Feoks/echo/pkg/task"
)

type sender struct {
	//TODO
}

// InitSender создает отправителя
func InitSender() Sender {
	return &sender{
		//TODO
	}
}

// Send отправляет данные
func (s sender) Send(tasks *[]task.Task) error {
	//TODO send to echo
	//for item := range *tasks {
	_, err := fmt.Println(tasks)
	if err != nil {
		return err
	}
	//}
	return nil
}

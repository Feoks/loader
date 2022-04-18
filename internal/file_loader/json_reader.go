package file_loader

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/Feoks/echo/pkg/task"
	"os"
)

type jsonReader struct {
	fileName string
	lastHash string
}

// InitJsonReader Создает и инициализирует новый json-ридер
func InitJsonReader(fName string) Reader {
	return &jsonReader{
		fileName: fName,
	}
}

// Read читает json-файл и возвращает срез тасков
func (jr *jsonReader) Read() (*[]task.Task, error) {
	jsonFile, err := os.ReadFile(jr.fileName)
	if err != nil {
		return nil, err
	}

	checksum := getFileHash(jsonFile)
	if checksum == jr.lastHash {
		return nil, nil
	}
	jr.lastHash = checksum

	var tasks *[]task.Task
	if err := json.Unmarshal(jsonFile, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// getFileHash Подсчитывает хэш файла
func getFileHash(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

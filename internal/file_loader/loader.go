package file_loader

type loader struct {
	reader Reader
	sender Sender
}

// initLoader создает экземпляр загрузчика
func initLoader(fileName string) Loader {
	return &loader{
		reader: InitJsonReader(fileName),
		sender: InitSender(),
	}
}

// Load Читает и отправляет данные
func (l loader) Load(fileName string) error {
	tasks, err := l.reader.Read()
	if err != nil {
		return err
	}
	if tasks == nil {
		return nil
	}
	err = l.sender.Send(tasks)
	return nil
}

package file_loader

type Loader interface {
	Load(fileName string) error
}

package main

type Reader interface {
	Read(data []byte) (int, error)
}

type Closer interface {
	Close() error
}

type ReadCloser interface {
	Reader
	Closer
}

type Opener func() (ReadCloser, error)

type File struct{}

func (f *File) Read(data []byte) (int, error) {
	return 0, nil
}

func (f *File) Close() error {
	return nil
}

func useFileProtocol(open Opener) {
	println("use file protocol")
	f, _ := open()
	data := make([]byte, 50)
	f.Read(data)
	f.Close()
}

func main() {
	useFileProtocol(func() (ReadCloser, error) {
		return &File{}, nil
	})
}

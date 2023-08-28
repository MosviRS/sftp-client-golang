package providers

import (
	"io"
	"os"
)

type TcpRepository interface {
	// Create a new file
	Create(path string) (io.ReadWriteCloser, error)
	Info(filePath string) (os.FileInfo, error)
}

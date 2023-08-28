package repositories

import (
	"fmt"
	"io"
	"os"
	"sftp-golang/src/client/sftp"
	"sftp-golang/src/errors"
	"sftp-golang/src/providers"
)

type SftpRepository struct {
	client *sftp.Client
}

func NewSftpRepository(client *sftp.Client) providers.TcpRepository {
	return &SftpRepository{client: client}
}

// Create creates a remote/destination file for I/O.
func (r *SftpRepository) Create(filePath string) (io.ReadWriteCloser, error) {
	if err := r.client.SftpConnection(); err != nil {
		return nil, errors.ErrorOpenConnection
	}
	return r.client.SftpClient.Create(filePath)
}

// Info gets the details of a file. If the file was not found, an error is returned.
func (r *SftpRepository) Info(filePath string) (os.FileInfo, error) {
	if err := r.client.SftpConnection(); err != nil {
		return nil, errors.ErrorOpenConnection
	}

	info, err := r.client.SftpClient.Lstat(filePath)
	if err != nil {
		return nil, fmt.Errorf("file stats: %w", err)
	}

	return info, nil
}

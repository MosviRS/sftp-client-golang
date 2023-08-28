package sftp

import (
	"fmt"
	"sftp-golang/src/errors"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Client provides basic functionality to interact with a SFTP server.
type Client struct {
	config     Config
	sshClient  *ssh.Client
	SftpClient *sftp.Client
}

// NewClient creates a new SFTP client.
func NewClient(config Config) (*Client, error) {
	client := &Client{
		config: config,
	}
	if client.SftpConnection() != nil {
		return nil, errors.ErrorOpenConnection
	}
	return client, nil
}

func (c *Client) isAlive() bool {
	_, _, err := c.sshClient.SendRequest(keepalive, false, nil)
	if err == nil {
		return true
	}
	return false
}

func (c *Client) SftpConnection() error {
	if c.sshClient != nil {
		if c.isAlive() {
			return nil
		}
	}

	auth := ssh.Password(c.config.Password)
	// if c.config.PrivateKey != "" {
	// 	signer, err := ssh.ParsePrivateKey([]byte(c.config.PrivateKey))
	// 	if err != nil {
	// 		return errors.ErrorParsePrivateKey
	// 	}
	// 	auth = ssh.PublicKeys(signer)
	// }

	cfg := &ssh.ClientConfig{
		User: c.config.Username,
		Auth: []ssh.AuthMethod{
			auth,
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         c.config.Timeout,
		Config: ssh.Config{
			KeyExchanges: c.config.KeyExchanges,
		},
	}

	sshClient, err := ssh.Dial("tcp", c.config.Server, cfg)
	if err != nil {
		return fmt.Errorf("ssh dial: %w", err)
	}
	c.sshClient = sshClient

	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		return fmt.Errorf("sftp new client: %w", err)
	}
	c.SftpClient = sftpClient

	return nil
}

// Close closes open connections.
func (c *Client) Close() {
	if c.SftpClient != nil {
		c.SftpClient.Close()
	}
	if c.sshClient != nil {
		c.sshClient.Close()
	}
}

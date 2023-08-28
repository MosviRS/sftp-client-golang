package sftp

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

var (
	keepalive = "keepalive"
	timeout   = time.Second * 30
)

// Config represents SSH connection parameters.
type Config struct {
	Username        string
	Password        string
	PrivateKey      string
	Server          string
	KeyExchanges    []string
	HostKeyCallback ssh.HostKeyCallback
	Timeout         time.Duration
}

func getHostKeyCallback() (ssh.HostKeyCallback, error) {
	hostKeyCallback, err := knownhosts.New(os.Getenv("SFTP_KNOWN_HOST_PATH"))
	if err != nil {
		return nil, err
	}
	return hostKeyCallback, nil
}
func getPrivateKey() (string, error) {
	key, err := ioutil.ReadFile(os.Getenv("SFTP_PRIVATE_KEY_PATH"))

	if err != nil {
		return "", err
	}
	return string(key), nil
}

func NewConfig(Username, Password, Server string) Config {
	hosts, err := getHostKeyCallback()
	if err != nil {
		log.Fatalf("Some error to get hosts: %s", err)
	}
	privateKey, err := getPrivateKey()
	if err != nil {
		log.Fatalf("Some error to get private key: %s", err)
	}
	return Config{
		Username:        Username,
		Password:        Password,
		PrivateKey:      privateKey,
		Server:          Server,
		HostKeyCallback: hosts,
		KeyExchanges:    []string{"diffie-hellman-group-exchange-sha256", "diffie-hellman-group14-sha256"},
		Timeout:         timeout,
	}
}

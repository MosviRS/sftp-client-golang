package main

import (
	"fmt"
	"log"
	"os"
	"sftp-golang/src/client/sftp"
	"sftp-golang/src/repositories"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured to read env. Err: %s", err)
	}
	config := sftp.NewConfig(os.Getenv("SFTP_USERNAME"), os.Getenv("SFTP_PASSWORD"), fmt.Sprintf("localhost:%s", os.Getenv("SFTP_PORT")))
	client, err := sftp.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	repoSftp := repositories.NewSftpRepository(client)

	// Get remote file stats.
	info, err := repoSftp.Info("./data")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", info)

	// Create remote file for writing.
	destination, err := repoSftp.Create("./data/file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer destination.Close()
	log.Println("Created file", destination)

}

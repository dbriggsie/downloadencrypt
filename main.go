package main

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/crypto/chacha20poly1305" 
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

const (
	aria2cPath        = "/usr/bin/aria2c"
	encryptedDirPath  = "./encrypted"
	publicKeyFilePath = "./publicKey.asc"
	minFreeSpace      = 5 * 1024 * 1024 * 1024 // 5GB
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Get the download URL from user input and validate it 
	fmt.Print("Enter download URL: ")
	url, _ := reader.ReadString('\n')
	url = strings.TrimSpace(url)
	if !strings.HasPrefix(url, "http") {
		log.Fatal("Invalid download URL")
	}

	// Check if the download file size is greater than or equal to the free space in the encrypted directory
	downloadFileInfo, err := os.Stat("./downloadedFile")
	if err != nil {
		log.Fatal(err)
	}
	if downloadFileInfo.Size() >= minFreeSpace {
		log.Fatal("Not enough free space in the encrypted directory")
	}

	// Download the file using Aria2c 
	// The split flag will allow aria2c to download the file using 16 connections, which can speed up the download process
	cmd := exec.Command(aria2cPath, "--split=16", url)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Check if there is enough free space in the encrypted directory for the downloaded file
	downloadedFileSize := downloadFileInfo.Size()
	encryptedDirStat, err := os.Stat(encryptedDirPath)
	if err != nil {
		log.Fatal(err)
	}
	if encryptedDirStat.Size()+downloadedFileSize >= minFreeSpace {
		log.Fatal("Not enough free space in the encrypted directory")
	}

	// Open the encrypted directory for writing
	encryptedDir, err := os.OpenFile(encryptedDirPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer encryptedDir.Close()

	// Create a new armored file for the encrypted data  
	armored, err := armor.Encode(encryptedDir, "PGP MESSAGE", map[string]string{
		"Version":           packet.DefaultVersion,
		"Comment":           "Encrypted data",
		"Hash":              "SHA256",
		"Charset":           "UTF-8",
		"Compression":       "ZIP",
		"Encryption":        "ChaCha20-Poly1305",
		"Encryption Cipher": "ChaCha20-Poly1305",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer armored.Close()

	// Get the recipient public key from user input
	fmt.Print("Enter recipient public key: ")
	publicKey, _ := reader.ReadString('\n')
	publicKey = strings.TrimSpace(publicKey)

	// Parse the recipient public key 
	entityList, err := openpgp.ReadArmoredKeyRing(strings.NewReader(publicKey))
	if err != nil {
		log.Fatal(err)
	}
	if len(entityList) < 1 {
		log.Fatal("No entities found in key ring")
	}
	entity := entityList[0]

	// Encrypt the file and write it to the armored file
	plainFile, err := os.Open("./downloadedFile")
	if err != nil {
		log.Fatal(err)
	}
	defer plainFile.Close()

	plainFileInfo, err := plainFile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	plaintext := make([]byte, plainFileInfo.Size())
	_, err = plainFile.Read(plaintext)
	if err != nil {
		log.Fatal(err)
	}

	writer, err := openpgp.Encrypt(armored, entityList, nil, &packet.Config{
		DefaultCipher: packet.CipherAES256,
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = writer.Write(plaintext)
	if err != nil {
		log.Fatal(err)
	}

	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Remove the plain file
	err = os.Remove("./downloadedFile")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File downloaded and encrypted successfully!")
}

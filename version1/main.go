package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/go-crypto/openpgp/armor"
	"github.com/ProtonMail/go-crypto/openpgp/packet"
)

// Path to the Aria2c executable
const aria2cPath = "/usr/bin/aria2c"

// Path to the encrypted directory
const encryptedDirPath = "./encrypted"

func main() {
	// Get the download URL from user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter download URL: ")
	url, _ := reader.ReadString('\n')
	url = strings.TrimSpace(url)

	// Download the file using Aria2c
	cmd := exec.Command(aria2cPath, url)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
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
		"Encryption":        "AES256",
		"Encryption Cipher": "AES256",
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

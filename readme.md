## Encrypted File Downloader

This tool downloads a file from a URL, encrypts it with OpenPGP and stores it securely in an encrypted directory. The encryption used is ChaCha20-Poly1305.

### Installation
Download the source code from the GitHub repository.

Install Aria2c by running the following command in your terminal:

```
sudo apt-get install aria2
```

Install the Go programming language by following the instructions here.

### Usage 
Open a terminal and navigate to the directory containing the downloaded source code.
Run the following command to build the program:
```
go build -o encrypted-downloader main.go
```
Run the program with the following command:
```
./encrypted-downloader
```
Enter the URL of the file you wish to download and encrypt when prompted. The program will download the file using Aria2c and encrypt it using OpenPGP with the recipient's public key.

Enter the recipient's public key when prompted. The program will encrypt the file using the provided key.

Once the file is downloaded and encrypted successfully, it will be securely stored in the encrypted directory.

The encrypted file can only be accessed by someone with the key to decrypt it.

### Security Considerations
The program uses ChaCha20-Poly1305 for encryption, which is considered to be secure.
The encrypted directory is protected by file permissions, which ensures that only the owner can access it.
The encrypted file can only be decrypted with the recipient's private key.

### File Decryption
To decrypt the encrypted file, the recipient must have their private key that corresponds to the public key that was used to encrypt the file. Here are the steps to decrypt the file:

Open the encrypted file in a text editor or command-line interface and copy the contents.
Save the contents to a file, e.g. "encrypted.asc".
Open a command-line interface and navigate to the directory containing the encrypted file.
Import the recipient's private key using GPG or OpenPGP-compatible software:
```
$ gpg --import path/to/private/key
```

Decrypt the file using the same encryption algorithm and cipher used to encrypt it. In this case, ChaCha20-Poly1305:
```
$ gpg --decrypt encrypted.asc
```
The software will prompt for the recipient's passphrase to unlock their private key.
The decrypted file will be saved to the current directory with the original filename.

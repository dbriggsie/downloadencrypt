# Encrypted File Downloader
This program downloads a file from a given URL and encrypts it using the recipient's public key. It uses the Aria2c downloader for fast and efficient downloads, and the OpenPGP library for encryption.

## Installation
### Prerequisites
- Go 1.16 or later
- Aria2c command line download utility
- GnuPG command line tool for generating a public key

## Install Aria2c
### Debian/Ubuntu
```
sudo apt-get install aria2
```
### macOS (using Homebrew)
```
brew install aria2
```

## Install GnuPG
### Debian/Ubuntu
```
sudo apt-get install gnupg
```
### macOS (using Homebrew)

```
brew install gnupg
```

## Build the DownloadEncrypt Program

Copy the source code into a file named *`main.go`*.
Open a terminal and navigate to the directory containing *`main.go`*.
Run the following command to build the executable:

```
go build -o encrypted-file-downloader main.go
```
This will create an executable named encrypted-file-downloader.

## Usage
Generate a public/private key pair
Open a terminal and run the following command to generate a new public/private key pair:
```
gpg --gen-key
```
Follow the prompts to create the key pair. Use the default settings for key type, key size, and expiration date.

Once the key pair is created, export the public key in ASCII armored format:

```
gpg --armor --export <your-email-address> > recipient_public_key.asc
```
This will save the public key to a file named recipient_public_key.asc.

Encrypt and download a file
Run the executable in a terminal:
```
./encrypted-file-downloader
```
Enter the download URL when prompted:
```
Enter download URL: <download URL>
```

Open the *`recipient_public_key.asc`* file and copy the entire content (including the header and footer).

Enter the recipient's public key in ASCII armored format when prompted:
```
Enter recipient public key: <paste the content of recipient_public_key.asc here>
```
The program will download the file, encrypt it using the recipient's public key, and save the encrypted file in the *`./encrypted directory`*.

Upon successful encryption and saving of the file, you will see the following message:

```
File downloaded and encrypted successfully!
```

*Note:* The program checks for available space in the ./encrypted directory before downloading and encrypting the file. If there is not enough space, the program will exit with an error message. Ensure that you have at least 5GB of free space in the ./encrypted directory before running the program.

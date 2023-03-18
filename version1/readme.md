Added constant values for the path to Aria2c executable and the encrypted directory path.
Changed the file open mode for the encrypted directory to `os.O_CREATE|os.O_WRONLY|os

Sure, here's an example code in Golang that you can compile. 
Please note that you'll need to have Aria2c installed on your system and have the relevant permissions to access the download folder.

You'll need to make sure that you have Aria2c and OpenSSL installed on your system. The code takes in a URL for a file to download, downloads it using Aria2c, and then encrypts it using OpenSSL with AES-256-CBC encryption. 
The encrypted file is then stored in a directory called encrypted, which is located in the same directory as the compiled binary.

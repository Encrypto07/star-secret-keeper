# Star Secret Keeper. ✨🔒

The Start Secret Keeper is a command-line interface (CLI) written in Go that allows you to securely manage secrets using AES encryption. It provides functionalities to set and retrieve secrets, ensuring that sensitive information is protected.

## Features 🚀

- Set a secret: Encrypts and stores a secret value securely.
- Get a secret: Retrieves and decrypts the stored secret value.

## Prerequisites 📋

- Go 1.16 or higher installed on your machine.

## Installation ⚙️

1. Clone the repository:

   ```shell
   git clone https://github.com/your-username/your-repo.git

## Steps to run the program
```shell
Steps:-

1-> go build -o secret                                                         Build the executable.

2-> ./secret set my-secret-value --key my-secret-keyxyz                        Set your secret.
    example -> ./secret set clientName::Adam --key JI0RUW255NF4PJJ0
        
3-> ./secret get --key my-secret-keyxyz                                        Get your secret.  
    example -> ./secret get --key JI0RUW255NF4PJJ0
    
Note-> Key length can be 8 or 16 🗝️
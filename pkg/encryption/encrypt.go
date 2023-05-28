package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func SetSecret(value string, key string) error {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	ciphertext := make([]byte, aes.BlockSize+len(value))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(value))

	if err := ioutil.WriteFile("secret.txt", ciphertext, 0644); err != nil {
		return err
	}

	fmt.Println("Secret set successfully!")
	return nil
}
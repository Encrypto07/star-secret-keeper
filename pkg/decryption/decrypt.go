package decryption

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"io/ioutil"
)

func GetSecret(key string) error {
	ciphertext, err := ioutil.ReadFile("secret.txt")
	if err != nil {
		return err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	if len(ciphertext) < aes.BlockSize {
		return errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	fmt.Println(string(ciphertext))
	return nil
}
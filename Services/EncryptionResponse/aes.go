package Encryptionresponse

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"os"
)

func Encrypt(text string) (string, error) {
	key := []byte(os.Getenv("AES_SECRET_KEY"))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, key[:block.BlockSize()])
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(cryptoText string) (string, error) {
	key := []byte(os.Getenv("AES_SECRET_KEY"))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext, err := hex.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBDecrypter(block, key[:block.BlockSize()])
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext), nil
}
package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func EncryptMessage(plainText, key string) (string, error) {
	keyBytes := []byte(key)
	plainTextBytes := []byte(plainText)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, plainTextBytes, nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func DecryptMessage(encryptedMessage, key string) (string, error) {
	keyBytes := []byte(key)
	cipherText, err := base64.StdEncoding.DecodeString(encryptedMessage)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]

	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

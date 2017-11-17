package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

func DecryptRSA(data string, privateKeyStr []byte) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	private, err := x509.ParsePKCS8PrivateKey(privateKeyStr)
	if err != nil {
		return "", err
	}
	privateKey := private.(*rsa.PrivateKey)
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)

	return string(decrypted), err
}

func DecryptBase64(data string) []byte {
	decrypted, _ := base64.StdEncoding.DecodeString(data)
	return decrypted
}

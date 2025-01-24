package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
)

func MD5(bytes []byte) string {
	h := md5.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}

func SHA1(bytes []byte) string {
	h := sha1.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}

func SHA256(bytes []byte) string {
	h := sha256.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}

var stringList = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(count int) string {
	b := make([]rune, count)
	for i := range b {
		b[i] = stringList[rand.Intn(len(stringList))]
	}
	return string(b)
}

// PKCS7Padding 用于填充
func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// PKCS7UnPadding 用于去除填充
func PKCS7UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])
	if unpadding > length {
		return nil, errors.New("invalid padding")
	}
	return src[:(length - unpadding)], nil
}

// DecryptECBWithClient 使用 AES-ECB 解密
func DecryptECBWithClient(cipherText string, key []byte) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	decrypted := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += aes.BlockSize {
		block.Decrypt(decrypted[i:i+aes.BlockSize], ciphertext[i:i+aes.BlockSize])
	}

	// 去除填充
	unpaddedText, err := PKCS7UnPadding(decrypted)
	if err != nil {
		return "", err
	}

	return string(unpaddedText), nil
}

func Encrypt(plainText string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := RandomString(gcm.NonceSize())
	encrypted := gcm.Seal([]byte(nonce), []byte(nonce), []byte(plainText), nil)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func Decrypt(cipherText string, key []byte) (string, error) {
	excepted, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(excepted) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, plainText := excepted[:nonceSize], excepted[nonceSize:]

	decrypted, err := gcm.Open(nil, nonce, plainText, nil)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

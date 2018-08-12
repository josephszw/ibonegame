package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func Bas(str string) {
	r := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(r)
}

// Aes128CBCEncrypt 加密 参数key为密钥，长度只能是16、24、32字节
func Aes128CBCEncrypt(str, serviceKey string) (ret string) {
	data := []byte(str)
	key, _ := base64.StdEncoding.DecodeString(serviceKey)
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	data = PKCS5Padding(data, blockSize)
	crypted := make([]byte, aes.BlockSize+len(data))
	iv := crypted[:aes.BlockSize]
	io.ReadFull(rand.Reader, iv)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(crypted[aes.BlockSize:], data)
	ret = base64.StdEncoding.EncodeToString(crypted)
	return
}

// Aes128CBCDecrypt 解密
func Aes128CBCDecrypt(data, serviceKey string, emptyIV bool) (ret string) {
	decryptText, _ := base64.StdEncoding.DecodeString(data)
	key, _ := base64.StdEncoding.DecodeString(serviceKey)
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	iv := make([]byte, blockSize)
	if emptyIV == false {
		iv = decryptText[:blockSize]
		decryptText = decryptText[blockSize:]
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(decryptText, decryptText)
	for i := len(decryptText) - 1; i > 0; i-- {
		if decryptText[i] != 0x00 {
			decryptText = decryptText[:i+1]
			break
		}
	}
	ret = string(decryptText)
	return
}

//PKCS5Padding Zeor Padding
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(0)}, padding)
	return append(ciphertext, padtext...)
}

//PKCS5UnPadding Zeor UnPadding
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

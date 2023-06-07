package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

const (
	AesKey = "8678f6b7bf2547b5aacd4f68074f6e2c"
	AesIv  = "2752084042949672"
)

type AesUtil struct {
	key      []byte
	keyBlock cipher.Block
	iv       []byte // 偏移向量，为了和第一组明文异或。解析见
}

func Constructor() (*AesUtil, error) {
	/*
		初始化结构体
	*/
	keyBlock, err := aes.NewCipher([]byte(AesKey))
	if err != nil {
		return nil, err
	} else {
		return &AesUtil{
			key:      []byte(AesKey),
			keyBlock: keyBlock,
			iv:       []byte(AesIv),
		}, nil
	}
}

func (AesUtil) Encrypt(content []byte) ([]byte, error) {
	/*
		CBC加密
		content 需要加密内容
	*/
	ca, err := Constructor()
	if err != nil {
		return nil, err
	}

	// 填充
	padText := PKCS7Padding(content, ca.keyBlock.BlockSize())
	blockMode := cipher.NewCBCEncrypter(ca.keyBlock, ca.iv)

	// 加密
	result := make([]byte, len(padText))
	blockMode.CryptBlocks(result, padText)
	// 返回密文
	return result, nil
}

func (AesUtil) Decrypt(encrypted []byte) ([]byte, error) {
	/*
		CBC 解密
		encrypted 待解密的密文
	*/
	ca, err := Constructor()
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(ca.keyBlock, ca.iv)
	result := make([]byte, len(encrypted))
	blockMode.CryptBlocks(result, encrypted)
	// 去除填充
	result = UnPKCS7Padding(result)
	return result, nil
}

func PKCS7Padding(text []byte, blockSize int) []byte {
	/*
		PKCS7Padding 填充模式
		text：明文内容
		blockSize：分组块大小
	*/
	padding := blockSize - len(text)%blockSize // 计算待填充的长度
	var paddingText []byte
	if padding == 0 {
		// 已对齐，填充一整块数据，每个数据为 blockSize
		paddingText = bytes.Repeat([]byte{byte(blockSize)}, blockSize)
	} else {
		// 未对齐 填充 padding 个数据，每个数据为 padding
		paddingText = bytes.Repeat([]byte{byte(padding)}, padding)
	}
	return append(text, paddingText...)
}

func UnPKCS7Padding(text []byte) []byte {
	/*
		去除 PKCS7Padding 填充的数据
		text 待去除填充数据的原文
	*/
	unPadding := int(text[len(text)-1]) // 取出填充的数据 以此来获得填充数据长度
	return text[:(len(text) - unPadding)]
}

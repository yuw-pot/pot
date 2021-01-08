// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package crypto

import (
	"bytes"
	cryptoAeS "crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	"reflect"
)

type (
	AeSPoT struct {
		key []byte
	}
)

func newAeS() *AeSPoT {
	return &AeSPoT {}
}

func (aes *AeSPoT) made(d ... interface{}) (interface{}, error) {
	if len(d) == 0 {
		return nil, E.Err(data.ErrPfx, "CryptoParamsErr")
	}

	if reflect.TypeOf(d[0]).Kind() != reflect.String {
		return nil, E.Err(data.ErrPfx, "CryptoAeSKeYErr")
	}

	simply := newSimply()
	h, err := simply.made(data.MD5, d[0])
	if err != nil {
		return nil, err
	}

	aes.key = []byte(cast.ToString(h)[8:24])
	return aes, nil
}

func (aes *AeSPoT) EncrypT(d interface{}) (interface{}, error) {
	enCode, err := aes.aesEncrypT([]byte(cast.ToString(d)))
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.EncodeToString(enCode), nil
}

func (aes *AeSPoT) DecrypT(cipherTxT interface{}) (interface{}, error) {
	deCode, err := base64.StdEncoding.DecodeString(cast.ToString(cipherTxT))
	if err != nil {
		return nil, err
	}

	d, err := aes.aesDecrypT(deCode)
	if err != nil {
		return nil, err
	}

	return string(d), nil
}

func (aes *AeSPoT) aesEncrypT(d []byte) ([]byte, error) {
	block, err := cryptoAeS.NewCipher(aes.key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	d = aes.pkcS5Padding(d, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, aes.key[:blockSize])
	cryptoDsT := make([]byte, len(d))
	blockMode.CryptBlocks(cryptoDsT, d)

	return cryptoDsT, nil
}

func (aes *AeSPoT) aesDecrypT(cryptoDsT []byte) ([]byte, error) {
	block, err := cryptoAeS.NewCipher(aes.key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, aes.key[:blockSize])
	d := make([]byte, len(cryptoDsT))
	blockMode.CryptBlocks(d, cryptoDsT)
	d = aes.pkcS5UnPadding(d)

	return d, nil
}

func (aes *AeSPoT) pkcS5Padding(cipherTxT []byte, blockSize int) []byte {
	padding := blockSize - len(cipherTxT)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherTxT, padText ...)
}

func (aes *AeSPoT) pkcS5UnPadding(d []byte) []byte {
	aesLength := len(d)
	unpadding := int(d[aesLength-1])
	return d[:(aesLength-unpadding)]
}





// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	F "github.com/yuw-pot/pot/modules/files"
)

type (
	RsAPoT struct {

	}
)

func newRsA() *RsAPoT {
	return &RsAPoT {

	}
}

func (rsa *RsAPoT) made(d ... interface{}) (interface{}, error) {
	return rsa, nil
}

func (rsa *RsAPoT) T() string {
	return "success T RsA!"
}


/**
 * Todo: Make PublicKey & PrivateKey
 *   - bits 1024, 2048 ...
 *   - Private Key File
 *   - Public Key File
**/
func RsA(bits int, filePrivateKey string, filePublicKey string) {
	var (
		err error
		fs *F.PoT = F.New()
	)

	// Make Private Key
	priKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}

	derStream := x509.MarshalPKCS1PrivateKey(priKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	priFile, errPriKey := fs.Create(filePrivateKey)
	if errPriKey != nil {
		panic(errPriKey)
	}

	defer func() {
		priFile.Close()
	}()

	if err = pem.Encode(priFile, block); err != nil {
		panic(err)
	}

	// Make Public Key
	pubKey := &priKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		panic(err)
	}

	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	pubFile, errPubKey := fs.Create(filePublicKey)
	if errPubKey != nil {
		panic(errPubKey)
	}

	defer func() {
		pubFile.Close()
	}()

	if err = pem.Encode(pubFile, block); err != nil {
		panic(err)
	}
}

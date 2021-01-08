// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	"github.com/yuw-pot/pot/modules/files"
)

const (
	defaultPath string = "./resources/pem/"
	defaultPriKeYName string = "privateKeY.pem"
	defaultPubKeYName string = "publicKeY.pem"
)

type (
	RsAPoT struct {
		pubKeY string
		priKeY string

		f *files.PoT
	}
)

func newRsA() *RsAPoT {
	return &RsAPoT {
		f: files.New(),
	}
}

func (rsaPoT *RsAPoT) made(d ... interface{}) (interface{}, error) {
	if len(d) > 3 {
		return nil, E.Err(data.ErrPfx, "CryptoParamsErr")
	}

	var (
		dirPath string = ""
		rsaPath string = defaultPath
	)

	if len(d) == 3 {
		dirPath = cast.ToString(d[2])
		if dirPath != "" {
			rsaPath = dirPath
		}
	}

	// d ... interface{} | d[0]:Public KeY, d[1]:Private KeY
	rsaPoT.pubKeY = rsaPath + cast.ToString(d[0])
	rsaPoT.priKeY = rsaPath + cast.ToString(d[1])

	return rsaPoT, nil
}

// encrypt by using Public Key
func (rsaPoT *RsAPoT) EncrypT(d string) (interface{}, error) {
	if rsaPoT.pubKeY == "" {
		return nil, E.Err(data.ErrPfx, "CryptoRsAPubKeY")
	}

	b, err := rsaPoT.getInfo(rsaPoT.pubKeY)
	if err != nil { return nil, err }

	pubKeY, err := x509.ParsePKIXPublicKey(b.Bytes)
	if err != nil { return nil, err }

	cipherTxT, err := rsa.EncryptPKCS1v15(rand.Reader, pubKeY.(*rsa.PublicKey), []byte(d))
	if err != nil { return nil, err }

	return base64.StdEncoding.EncodeToString(cipherTxT), nil
}

// decrypt by using private Key
func (rsaPoT *RsAPoT) DecrypT(cipherTxT interface{}) (interface{}, error) {
	if rsaPoT.priKeY == "" {
		return nil, E.Err(data.ErrPfx, "CryptoRsAPriKeY")
	}

	byteCipherTxT, err := base64.StdEncoding.DecodeString(cast.ToString(cipherTxT))
	if err != nil { return nil, err }

	b, err := rsaPoT.getInfo(rsaPoT.priKeY)
	if err != nil { return nil, err }

	priKeY, err := x509.ParsePKCS1PrivateKey(b.Bytes)
	if err != nil { return nil, err }

	d, err := rsa.DecryptPKCS1v15(rand.Reader, priKeY, byteCipherTxT)
	if err != nil { return nil, err }

	return cast.ToString(d), err
}

func (rsaPoT *RsAPoT) getInfo(keyPath string) (*pem.Block, error) {
	var err error

	fp, err := rsaPoT.f.Open(keyPath)
	if err != nil { return nil, err }

	defer fp.Close()

	key, err := fp.Stat()
	if err != nil { return nil, err }

	buf := make([]byte, key.Size())
	_, err = fp.Read(buf)
	if err != nil { return nil, err }

	block, _ := pem.Decode(buf)
	return block, nil
}

/**
 * Todo: Make PublicKey & PrivateKey
 *   - bits 1024, 2048 ...
 *   - Private Key File
 *   - Public Key File
**/
func RsA(bits int) {
	var (
		err error
		fs *files.PoT = files.New()
	)

	// Make Private Key
	priKeY, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}

	derStream := x509.MarshalPKCS1PrivateKey(priKeY)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	err = fs.MkdirAll(defaultPath)
	if err != nil {
		panic(err)
	}

	priFile, errPriKey := fs.Create(defaultPath + defaultPriKeYName)
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
	pubKeY := &priKeY.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(pubKeY)

	if err != nil {
		panic(err)
	}

	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	err = fs.MkdirAll(defaultPath)
	if err != nil {
		panic(err)
	}

	pubFile, errPubKey := fs.Create(defaultPath + defaultPubKeYName)
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

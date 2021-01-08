// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package files

import (
	"os"
)

type PoT struct {}

func New() *PoT {
	return &PoT {}
}

func (fs *PoT) IsExists(pathname string) (bool, error) {
	var (
		ok bool = false
		err error
	)

	_, err = os.Stat(pathname)
	if err == nil {
		ok = true
		return ok, err
	}

	if os.IsNotExist(err) {
		ok = false
	} else {
		ok = true
	}

	return ok, err
}

func (fs *PoT) Open(pathname string) (*os.File, error) {
	return os.Open(pathname)
}

func (fs *PoT) Create(pathname string) (*os.File, error) {
	return os.Create(pathname)
}

func (fs *PoT) Mkdir(dir string) error {
	ok, _ := fs.IsExists(dir)
	if ok {
		return nil
	}

	return os.Mkdir(dir, os.ModePerm)
}

func (fs *PoT) MkdirAll(dir string) error {
	ok, _ := fs.IsExists(dir)
	if ok {
		return nil
	}

	return os.MkdirAll(dir, os.ModePerm)
}

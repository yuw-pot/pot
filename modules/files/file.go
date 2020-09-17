// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package files

import "os"

type (
	PoT struct {

	}
)

func New() *PoT {
	return &PoT {

	}
}

func (fs *PoT) IsExists(pathname string) (ok bool, err error) {
	_, err = os.Stat(pathname)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, err
	}

	return
}

func (fs *PoT) Open(pathname string) (f *os.File, err error) {
	f, err = os.Open(pathname)
	defer f.Close()

	return
}

func (fs *PoT) Create(pathname string) (f *os.File, err error) {
	f, err = os.Create(pathname)
	defer f.Close()

	return
}

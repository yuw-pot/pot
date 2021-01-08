// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package txt

import (
	"bufio"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/modules/files"
	"github.com/yuw-pot/pot/modules/properties"
	"io/ioutil"
	"strings"
)

const defaultPath = "./resources/docs"

type Component struct {
	f *files.PoT
	err error
	content []byte

	dir string
}

func New() *Component {
	dir := properties.PropertyPoT.GeT("TxT.Path", defaultPath)
	return &Component {
		f: files.New(),
		dir: cast.ToString(dir),
	}
}

func (txt *Component) ReadAll(filename string) *Component {
	dir := strings.Join([]string{txt.dir, filename}, "/")
	_, err := txt.f.IsExists(dir)
	if err != nil {
		txt.err = err
		return txt
	}

	fp, err := txt.f.Open(dir)
	if err != nil {
		txt.err = err
		return txt
	}

	defer fp.Close()

	r, err := ioutil.ReadAll(fp)
	if err != nil {
		txt.err = err
		return txt
	}

	txt.content = r
	return txt
}

func (txt *Component) ReadLine(start, limit int, filename string) ([]string, error) {
	dir := strings.Join([]string{txt.dir, filename}, "/")
	_, err := txt.f.IsExists(dir)
	if err != nil {
		return nil, err
	}

	fp, err := txt.f.Open(dir)
	if err != nil {
		return nil, err
	}

	defer fp.Close()

	var (
		fs *bufio.Scanner = bufio.NewScanner(fp)
		i int = 0
		num int = 1
		content []string = make([]string, limit)
	)

	for fs.Scan() {
		if num >= start && num < start+limit {
			content[i] = fs.Text()
			i ++
		}

		num ++
	}

	return content, nil
}

func (txt *Component) GeTByLine() {
	// todo: do something
}

func (txt *Component) Byte() []byte {
	return txt.content
}

func (txt *Component) String() string {
	return string(txt.content)
}

func (txt *Component) Error() error {
	return txt.err
}

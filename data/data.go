// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package data

const (
	PoT	string = "PoT"

	// PoT Mode
	ReleaseMode	string = "release"
	ConsoleMode string = "console"
)

var (
	PoTMode []interface{} = []interface{}{
		ConsoleMode,
		ReleaseMode,
	}
)

type (
	// Default Type
	H map[string]interface{}

	// Err Type
	ErrH map[string]H

	// Power
	PowerPoT struct {
		JwT 	int `yaml:"jwt" json:"jwt"`
		Adapter int	`yaml:"adapter" json:"adapter"`
		Redis 	int `yaml:"redis" json:"redis"`
	}

	TpL struct {
		Status int
		Msg interface{}
		Response *H
	}

	SrvTpL struct {
		Status int
		Msg interface{}
		Data *H
	}
)

func SrvTpLInitialized() *SrvTpL {
	return &SrvTpL {
		Status: PoTStatusOK,
		Msg:    nil,
		Data:   nil,
	}
}

func TpLInitialized() *TpL {
	return &TpL {
		Status: PoTStatusOK,
		Msg:    nil,
		Response: &H{},
	}
}

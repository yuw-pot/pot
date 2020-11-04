// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package data

const (
	// Models
	ModONE		string = "ONE"
	ModALL		string = "ALL"
	ModByAsc	string = "ASC"
	ModByEsc	string = "DESC"

	// ModJoinPoT
	ModJoinLEFT		string = "LEFT"
	ModJoinINNER	string = "INNER"
	ModJoinRIGHT	string = "RIGHT"
)

type (
	// Models Struct
	ModPoT struct {
		Types 		string `json:"types"`// Todo: "ONE"->"FetchOne" "ALL"->"FetchAll"
		Table 		string `json:"table"`
		Field 		[]string `json:"field"`
		Joins 		[]*ModJoinPoT `json:"joins"`
		Limit 		int `json:"limit"`
		Start 		[]int `json:"start"`
		Query 		interface{} `json:"query"`
		QueryArgs 	[]interface{} `json:"query_args"`
		Columns 	[]string `json:"columns"`
		OrderType 	string `json:"order_type"`
		OrderArgs 	[]string `json:"order_args"`
	}

	ModJoinPoT struct {
		JoinOperator string
		TableName interface{}
		Condition string
	}
)

// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package data

const (
	ReqMethodGET	= "GET"
	ReqMethodPUT	= "PUT"
	ReqMethodPOST	= "POST"
)

var (
	ReqMethods []interface{} = []interface{}{
		ReqMethodGET,
		ReqMethodPUT,
		ReqMethodPOST,
	}
)

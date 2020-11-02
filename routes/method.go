// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

const (
	PoTMethodAnY	 = "ANY"
	PoTMethodGeT  	 = "GET"
	PoTMethodHeaD    = "HEAD"
	PoTMethodPosT    = "POST"
	PoTMethodPuT     = "PUT"
	PoTMethodPatch   = "PATCH" // RFC 5789
	PoTMethodDelete  = "DELETE"
	PoTMethodConnect = "CONNECT"
	PoTMethodOptions = "OPTIONS"
	PoTMethodTrace   = "TRACE"
)

func GeTPath(service string, controller string, action string) interface{} {
	_, ok := (*rMaP)[service + controller + action]
	if ok == false {
		return nil
	}

	return (*rMaP)[service + controller + action]
}

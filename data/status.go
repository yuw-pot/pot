// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package data

import (
	"net/http"
)

const (
	// Status Success
	PoTStatusOK			= http.StatusOK

	// Status Other
	PoTUnKnown			= -1
	PoTStatusNotFound	= http.StatusNotFound
	PoTStatusNoContent	= http.StatusNoContent
)

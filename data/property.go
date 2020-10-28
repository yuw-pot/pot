// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package data

const (
	// Property File Suffix
	EnvDEV 		string = "dev"
	EnvSTG 		string = "stg"
	EnvPRD 		string = "prd"
	PropertySfx string = "yaml"

	PropertyPort 			string = "8577"
	PropertyMode 			int = 0
	PropertyTimeLocation 	string = "Asia/Shanghai"
	PropertyHsslPower		int = 0

	PropertyJwT				int = 1
	PropertyAdapter			int = 1
	PropertyRedis			int = 1
)

var (
	/**
	 * Modules: Property
	 *   - dev
	 *   - stg
	 *   - prd
	 */
	PropertySfxs []interface{} = []interface{}{
		EnvDEV,
		EnvSTG,
		EnvPRD,
	}
)

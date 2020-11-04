// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package services

import (
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/routes"
	"math"
)

const (
	defaultPage int = 1
	defaultSize int = 10
)

type (
	Services struct {

	}
)

func New() *Services {
	return &Services {

	}
}

func (srv *Services) GeTPath(service, controller, action string) (interface{}, error) {
	return routes.GeTPath(service, controller, action)
}

func (srv *Services) Paginator(page int, pageNums int, pageSize ...int) (paginator map[string]interface{}) {
	var (
		nums int = pageNums
		size int
		prePage int
		sufPage int
	)

	if len(pageSize) > 0 {
		size = pageSize[0]
	} else {
		size = defaultSize
	}

	var totalPage int = int(math.Ceil(float64(nums) / float64(size)))

	if page > totalPage {
		page = totalPage
	}

	if page <= 0 {
		page = 1
	}

	var pages []int

	switch {
	case page >= totalPage-5 && totalPage > 5:
		start := totalPage-5+1
		prePage = page-1
		sufPage = int(math.Min(float64(totalPage), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start+i
		}

	case page >= 3 && totalPage > 5:
		start := page-3+1
		pages = make([]int, 5)
		prePage = page-3
		for i, _ := range pages {
			pages[i] = start + i
		}

		prePage = page-1
		sufPage = page+1

	default:
		pages = make([]int, int(math.Min(5, float64(totalPage))))
		for i, _ := range pages {
			pages[i] = i + 1
		}

		prePage = int(math.Max(float64(1), float64(page-1)))
		sufPage = page+1
	}

	paginator = map[string]interface{}{
		"pages": pages,
		"total": totalPage,
		"cur_page": page,
		"pre_page": prePage,
		"suf_page": sufPage,
	}

	return
}

func (srv *Services) PaginatorParams(strPage string, strSize string) (page int, size int) {
	if strPage != "" {
		page = cast.ToInt(strPage)
	} else {
		page = defaultPage
	}

	if strSize != "" {
		size = cast.ToInt(strSize)
	} else {
		size = defaultSize
	}

	return
}

package helper

import (
	"fmt"
	"strings"
)

// JqGrid ...
type JqGrid struct {
	Pagenum  int `query:"pagenum"`
	Pagesize int `query:"pagesize"`

	Sortfield string `query:"sortfield"`
	Sortorder string `query:"sortorder"`

	Filterscount    int      `query:"filterscount"`
	Filterdatafield []string `query:"filterdatafield"`
	Filtervalue     []string `query:"filtervalue"`
}

// SelectSortFilterData ...
func SelectSortFilterData(p *JqGrid, allFilter string) (offset, limit, filterBy, orderBy string) {
	pageNo := p.Pagenum
	pageSize := p.Pagesize

	//fmt.Printf("pagenum=%d and pagesize=%d", pageNo, pageSize)

	if pageNo < 0 {
		pageNo = 0
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset = fmt.Sprintf("%d", pageNo*pageSize)
	limit = fmt.Sprintf("%d", pageSize)

	filterBy = ""
	if len(p.Filterdatafield) > 0 && strings.ToUpper(p.Filterdatafield[0]) == "ALL" {
		filterBy = allFilter + " LIKE UPPER('%" + p.Filtervalue[0] + "%')"
	} else {
		for i := 0; i < p.Filterscount; i++ {
			if p.Filtervalue[i] != "null" {
				filterBy += " AND UPPER(CAST(" + p.Filterdatafield[i] + " AS TEXT)) LIKE UPPER('%" + p.Filtervalue[i] + "%')"
			}
		}
	}
	/*if p.Filterfield != "" && p.Filtertext != "" {
		filterBy += " AND UPPER(CAST(" + p.Filterfield + " AS TEXT)) LIKE UPPER('%" + p.Filtertext + "%')"
	}*/

	if p.Sortfield == "" {
		p.Sortfield = "id"
	}
	orderBy = p.Sortfield + " " + p.Sortorder

	return
}

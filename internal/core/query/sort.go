package query

import (
	"strings"
	"github.com/go-jet/jet/v2/qrm"
)

const (
	randomSortKey = "@random"
	rowidSortKey  = "@rowid"
)

const (
	SortAsc  = "ASC"
	SortDesc = "DESC"
)

type SortFiled struct {
	Name      string `json:"name"`
	Direction string `json:"direction"`
}

func ParseSortFromString(str string) (fields []SortFiled) {
	data := strings.Split(str, ",")

	for _, field := range data {
		field = strings.TrimSpace(field)
		if field == "" {
			continue
		}
		switch {
		case strings.HasPrefix(field, "-"):
			fields = append(fields, SortFiled{strings.TrimPrefix(field, "-"), SortDesc})
		case strings.HasPrefix(field, "+"):
			fields = append(fields, SortFiled{strings.TrimPrefix(field, "+"), SortAsc})
		default:
			fields = append(fields, SortFiled{field, SortAsc})
		}
	}
	return
}

type ColumnResolver func(field string) (qrm, error)

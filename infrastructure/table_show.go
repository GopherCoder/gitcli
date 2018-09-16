package infrastructure

import (
	"fmt"
	"strings"

	"github.com/alexeyco/simpletable"
)

func TableShow(values []interface{}, fields []string) {
	table := simpletable.New()
	setHeader(fields, table)
	setValue(values, table)
	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())

}

func setHeader(fields []string, table *simpletable.Table) *simpletable.Table {
	cells := []*simpletable.Cell{}
	for _, field := range fields {
		cells = []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: strings.ToUpper(field)},
		}
	}
	table.Header = &simpletable.Header{
		Cells: cells,
	}
	return table

}

func setValue(values []interface{}, table *simpletable.Table) *simpletable.Table {
	for _, value := range values {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: value.(string)},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}
	return table

}

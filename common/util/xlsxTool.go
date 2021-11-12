package util

import (
	"errors"

	"github.com/tealeg/xlsx/v3"
)

type XlsxRow struct {
	Row  *xlsx.Row
	Data []string
}

func NewRow(row *xlsx.Row, data []string) *XlsxRow {
	return &XlsxRow{
		Row:  row,
		Data: data,
	}
}

func (row *XlsxRow) SetRowTitle() error {
	return generateRow(row.Row, row.Data)
}

func (row *XlsxRow) GenerateRow() error {
	return generateRow(row.Row, row.Data)
}

func generateRow(row *xlsx.Row, rowStr []string) error {
	if rowStr == nil {
		return errors.New("no data to generate xlsx!")
	}
	for _, v := range rowStr {
		cell := row.AddCell()
		cell.SetString(v)
	}
	return nil
}

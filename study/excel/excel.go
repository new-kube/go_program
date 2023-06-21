package main

import (
	"fmt"

	xlsx "github.com/tealeg/xlsx/v3"
)

type Data struct {
	Code  string
	Name  string
	Count int
}

func operate(in string, out string) {
	inFile, err := xlsx.OpenFile(in)
	if err != nil {
		fmt.Printf("open file=%s err: %v\n", in, err)
		return
	}

	if len(inFile.Sheets) == 0 {
		fmt.Printf("no sheets in %s\n", in)
		return
	}

	sheet := inFile.Sheets[0]

	outFile := xlsx.NewFile()
	oSheet, err := outFile.AddSheet(sheet.Name)
	if err != nil {
		fmt.Printf("add sheet=%s err: %v\n", sheet.Name, err)
		return
	}

	oRow := oSheet.AddRow()
	for i := 0; i < 3; i++ {
		cell, err := sheet.Cell(0, i)
		if err != nil {
			fmt.Printf("get cell=(0,%d) err: %v\n", i, err)
			return
		}
		oCell := oRow.AddCell()
		oCell.SetString(cell.String())
		oCell.SetStyle(cell.GetStyle())
		fmt.Printf("cell=%+v\n", *cell)
		fmt.Printf("cell value=%s, string=%s\n", cell.Value, cell.String())
	}

	d1 := Data{
		Code:  "12345",
		Name:  "第一本书",
		Count: 100,
	}

	row := oSheet.AddRow()

	row.AddCell().SetString(d1.Code)
	row.AddCell().SetString(d1.Name)
	row.AddCell().SetInt(d1.Count)

	d2 := Data{
		Code:  "123456",
		Name:  "第二本书",
		Count: 200,
	}

	row = oSheet.AddRow()

	row.AddCell().SetString(d2.Code)
	row.AddCell().SetString(d2.Name)
	row.AddCell().SetInt(d2.Count)

	if err := outFile.SaveAs(out); err != nil {
		fmt.Printf("save file=%s err: %v\n", out, err)
		return
	}
}

func main() {
	in, out := "in.xlsx", "out.xlsx"
	operate(in, out)
}

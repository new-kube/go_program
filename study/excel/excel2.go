package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Data struct {
	Code  string
	Name  string
	Count int
}

/*
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
*/

func main() {
	in, out := "in.xlsx", "out.xlsx"
	test(in, out)
}

func test(in, out string) {
	f, err := excelize.OpenFile(in)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheets := f.GetSheetMap()
	names := make([]string, 0, len(sheets))
	for _, name := range sheets {
		names = append(names, name)
	}

	sheet := names[0]

	d1 := Data{
		Code:  "12345",
		Name:  "第一本书",
		Count: 100,
	}

	x1 := &[]interface{}{d1.Code, d1.Name, d1.Count}
	if err := f.SetSheetRow(sheet, "A2", x1); err != nil {
		fmt.Println(err)
		return
	}

	d2 := Data{
		Code:  "123456",
		Name:  "第二本书",
		Count: 200,
	}

	x2 := &[]interface{}{d2.Code, d2.Name, d2.Count}
	if err := f.SetSheetRow(sheet, "A3", x2); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SaveAs(out); err != nil {
		fmt.Println(err)
		return
	}
}

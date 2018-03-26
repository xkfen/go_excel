package service

import (
	"github.com/tealeg/xlsx"
	"fmt"
)

// 创建excel
func CreateTealegXlsx(){
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println(err.Error())
	}
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "test"
	err = file.Save("/home/qvdev/test.xlsx")
	if err != nil {
		fmt.Println(err.Error())
	}
}


// 读取excel

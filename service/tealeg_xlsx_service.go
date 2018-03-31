package service

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"github.com/novalagung/gubrak"
)
	/**
	该service引用的xlsx库为github.com/tealeg/xlsx
	 */

// 创建excel
func CreateTealegXlsx(){
	// 创建一个xlsx文件
	file := xlsx.NewFile()
	// Add a new Sheet, with the provided name, to a File
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println(err.Error())
	}
	// Add a new Row to a Sheet
	row := sheet.AddRow()
	// Add a new Cell to a Row
	cell := row.AddCell()
	cell.Value = "test"
	//err = file.Save("./test.xlsx")
	// Save the File to an xlsx file at the provided path.
	err = file.Save("/home/qydev/test.xlsx")
	if err != nil {
		fmt.Println(err.Error())
	}
}


// 读取excel
func ReadTealegXlsx(){
	excelFileName := "/home/qydev/test.xlsx"
	// 根据指定的xlsx名称，返回xlsx文件结构
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	// xlFile.Sheets :一个xlsx文件可能有多个Sheet
	for _, sheet := range xlFile.Sheets {
		// sheet.Rows:得到sheet的每一行
		for _, row := range sheet.Rows {
			// row.Cells:每一行的每一列
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}


func Chunk(){
	data := []int{ 1, 2, 3, 4, 5 }
	size := 2
	info, err := gubrak.Chunk(data, size)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(info)
}
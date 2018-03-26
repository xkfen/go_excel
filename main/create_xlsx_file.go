package main

import(
	"github.com/360EntSecGroup-Skylar/excelize"
	"fmt"
)
func main(){
	// 使用默认模板创建一个新的文件
	xlsx := excelize.NewFile()
	// 默认的excel模板第一个sheet为Sheet1，这里将默认name修改为测试
	xlsx.SetSheetName("Sheet1", "测试")
	// 得到该sheet的索引
	index := xlsx.GetSheetIndex("测试")
	xlsx.SetActiveSheet(index)
	// Create a new sheet.
	//index := xlsx.NewSheet("Sheet1")
	// set title
	xlsx.SetCellStr("Sheet1", "A1", "期数")
	xlsx.SetSheetBackground("Sheet1", "/home/qydev/id_card_img.jpeg")
	// Set value of a cell.
	xlsx.SetCellValue("Sheet1", "A2", "Hello world.")
	// Set active sheet of the workbook.
	//xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./测试.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
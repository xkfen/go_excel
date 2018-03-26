package service

import(
	"github.com/360EntSecGroup-Skylar/excelize"
	"fmt"
)
// 引用的这个package里，只有file, sheet,cell,没有row这个概念
func CreateExcel(){
	// 创建excel
	xlsx := excelize.NewFile()
	// 给创建的sheet设置名字
	xlsx.SetSheetName("Sheet1", "测试")
	// 得到当前sheet的index
	index := xlsx.GetSheetIndex("测试")
	// 得到当前excel的name
	//xlsx.GetSheetName(index)
	// 将当前sheet的设置为激活状态
	xlsx.SetActiveSheet(index)
	// 给A1列设置值:这里明确设置为str
	xlsx.SetCellStr("测试", "A1","标题1")
	// 如果不知道数据类型，可以使用SetCellValue
	xlsx.SetCellValue("测试","A2", 100)
	xlsx.SetCellStr("测试", "B1","标题2")
	xlsx.SetCellStr("测试", "B2","XXX")
	xlsx.SetCellStr("测试", "C1","标题3")
	xlsx.SetCellStr("测试", "C2","CCC")
	// 得到该excel中活跃的sheet index
	//n := xlsx.GetActiveSheetIndex()
	// get worksheet name and index map of XLSX:key为sheet的index，value为sheetName
	xlsxIntMap := xlsx.GetSheetMap()
	fmt.Println(xlsxIntMap)
	// todo 为给定的sheet设置背景图,这个也没有效果
	xlsx.SetSheetBackground("测试", "/home/qydev/bg1.jpg")
	// 新建一个Sheet
	xlsx.NewSheet("Sheet2")
	xlsx.NewSheet("Sheet3")
	// 删掉sheet，如果有与之关联的sheet会受到影响
	//xlsx.DeleteSheet("Sheet2")
	// 复制,from为源sheet index,to 为目的sheet index
	//xlsx.CopySheet(index, 2)
	// 设置sheet是否可见，一个excel必须有一个sheet可见，如果当前的sheet是activated的，那么该方法无效
	// todo 这个方法有问题，无法设置sheet2为false，可以设置sheet3
	//xlsx.SetSheetVisible("Sheet3", false)
	// 保存excel
	err := xlsx.SaveAs("./测试.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func ReadExcel(){
	xlsx , err := excelize.OpenFile("测试.xlsx")
	if err != nil {
		fmt.Println("读取文件出错", err.Error())
	}

	// 得到特定的某列的值
	a1 := xlsx.GetCellValue("测试","A1")
	fmt.Println(a1)

	// 得到指定sheet中所有的行
	rows := xlsx.GetRows("测试")
	// 循环遍历每一行，
	for _, row := range rows{
		for _, cell := range row{
			fmt.Print(cell, "\t")
		}
		fmt.Println()
	}
}
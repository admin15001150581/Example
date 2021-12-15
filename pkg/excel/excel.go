package excel

import (
	"github.com/xuri/excelize/v2"
	"example/models"
	"strconv"
)
//excel导出
func ExportExcel(tag []models.Tag) error{
	f := excelize.NewFile()
	// 创建一个工作表
	//index := f.NewSheet()
	//设置表头
	f.SetCellValue("Sheet1", "A1","id" )
	f.SetCellValue("Sheet1", "B1","名字" )
	f.SetCellValue("Sheet1", "C1","创建时间" )
	f.SetCellValue("Sheet1", "D1","创建人" )
	f.SetCellValue("Sheet1", "E1","修改时间" )
	f.SetCellValue("Sheet1", "F1","修改人" )
	f.SetCellValue("Sheet1", "G1","删除时间" )
	f.SetCellValue("Sheet1", "H1","状态值" )
	// 设置单元格的值
	for key,value:=range tag{
		zhi:=key+2
		string:=strconv.Itoa(zhi) //把整型转换为字符串类型
		f.SetCellValue("Sheet1", "A"+string, value.ID)
		f.SetCellValue("Sheet1", "B"+string, value.Name)
		f.SetCellValue("Sheet1", "C"+string, value.CreatedOn)
		f.SetCellValue("Sheet1", "D"+string, value.CreatedBy)
		f.SetCellValue("Sheet1", "E"+string, value.ModifiedOn)
		f.SetCellValue("Sheet1", "F"+string, value.ModifiedBy)
		f.SetCellValue("Sheet1", "G"+string, value.DeletedOn)
		f.SetCellValue("Sheet1", "H"+string, value.State)
	}

	//f.SetCellValue("Sheet2", "A2", "Hello world.")
	//f.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	//f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		return err
	}
	return  nil
}

package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func ExchangeExcel() {

	xlsx1, err1 := excelize.OpenFile("./document/Book1.xlsx")

	if err1 != nil {
		fmt.Println(err1)
		return
	}

	xlsx2, err2 := excelize.OpenFile("./document/Book2.xlsx")

	if err2 != nil {
		fmt.Println(err2)
		return
	}

	rows := xlsx1.GetRows("Sheet1")

	for i, row := range rows {

		for j, colCell := range row {

			rowa_str := "A" + strconv.Itoa(i*7+j+1)
			rowb_str := "B" + strconv.Itoa(i*7+j+1)

			colaVal := ""
			if j == 0 {
				colaVal = "日期:"
			} else if j == 1 {
				colaVal = "客户名称:"
			} else if j == 2 {
				colaVal = "型材:"
			} else if j == 3 {
				colaVal = "高:"
			} else if j == 4 {
				colaVal = "宽:"
			} else if j == 5 {
				colaVal = "玻璃款式:"
			} else if j == 6 {
				colaVal = "备注:"
			}

			xlsx2.SetCellValue("Sheet1", rowa_str, colaVal)
			//fmt.Print(colCell)

			if j == 0 {
				if strings.Contains(colCell, "-") {
					date_str := ""
					date_arry := strings.Split(colCell, "-")
					date_str = date_arry[2] + "-" + date_arry[0] + "-" + date_arry[1]
					xlsx2.SetCellValue("Sheet1", rowb_str, date_str)

				} else {
					xlsx2.SetCellValue("Sheet1", rowb_str, colCell)
				}

			} else if j == 3 {
				f, err := strconv.ParseFloat(colCell, 64)
				if err != nil {
					fmt.Println(err)
				}
				xlsx2.SetCellValue("Sheet1", rowb_str, f)

			} else if j == 4 {
				f, err := strconv.ParseFloat(colCell, 64)
				if err != nil {
					fmt.Println(err)
				}
				xlsx2.SetCellValue("Sheet1", rowb_str, f)

			} else {
				xlsx2.SetCellValue("Sheet1", rowb_str, colCell)
			}

		}

	}

	//xlsx2.SetActiveSheet(index2)

	err3 := xlsx2.SaveAs("./document/Book2.xlsx")
	if err3 != nil {
		fmt.Println(err3)
	}

}

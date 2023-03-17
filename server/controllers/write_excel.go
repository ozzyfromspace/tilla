package controllers

import (
	"fmt"
	"log"
	"tilla/models"
	"time"

	"github.com/xuri/excelize/v2"
)

type ConversionDoc struct {
	Student models.Student
	Events  models.Events
}

func generateExcel(convDoc *[]ConversionDoc, startDay int, startMonth time.Month, startYear int, endDay int, endMonth time.Month, endYear int) (string, error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	log.Printf("GENERATING EXCELS! %+v, %+v\n", (*convDoc)[0].Student, (*convDoc)[0].Events)

	for ci, el := range *convDoc {
		events := el.Events
		student := el.Student

		if len(events) == 0 {
			continue
		}

		sheetName := fmt.Sprintf("%v %v - %v", student.FirstName, student.LastName, truncateStr(student.Id))
		index, err := f.NewSheet(sheetName)
		if err != nil {
			fmt.Println(err)
			return "", err
		}

		f.SetCellValue(sheetName, "A1", "Course")
		f.SetCellValue(sheetName, "B1", "Teacher")
		f.SetCellValue(sheetName, "C1", "Date")
		f.SetCellValue(sheetName, "D1", "Time")
		f.SetCellValue(sheetName, "E1", "Duration")
		f.SetCellValue(sheetName, "F1", "Fee")
		f.SetCellValue(sheetName, "G1", "Payment Status")
		f.SetCellValue(sheetName, "H1", "Rate ($/Session)")
		f.SetCellValue(sheetName, "I1", "Session Length")

		for i, event := range events {
			f.SetCellValue(sheetName, coordinate(0, i+1), event.Course)
			f.SetCellValue(sheetName, coordinate(1, i+1), event.Teacher)
			f.SetCellValue(sheetName, coordinate(2, i+1), event.Date)
			f.SetCellValue(sheetName, coordinate(3, i+1), event.Time)
			f.SetCellValue(sheetName, coordinate(4, i+1), event.Duration)
			f.SetCellValue(sheetName, coordinate(5, i+1), event.Fee)
			f.SetCellValue(sheetName, coordinate(6, i+1), event.PaymentStatus)
			f.SetCellValue(sheetName, coordinate(7, i+1), event.Rate)
			f.SetCellValue(sheetName, coordinate(8, i+1), event.SessionLengthInMinutes)
		}

		formulaSum := fmt.Sprintf("=SUM(F2:F%d)", len(events)+1)
		formulaAverage := fmt.Sprintf("=AVERAGE(F2:F%d)", len(events)+1)

		f.SetCellValue(sheetName, fmt.Sprintf("E%d", len(events)+3), "Total")
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", len(events)+4), "Average")

		f.SetCellFormula(sheetName, fmt.Sprintf("F%d", len(events)+3), formulaSum)
		f.SetCellFormula(sheetName, fmt.Sprintf("F%d", len(events)+4), formulaAverage)

		if ci == 0 {
			f.SetActiveSheet(index)
			if err := f.DeleteSheet("Sheet1"); err != nil {
				log.Println(err)
			}
		}
	}

	filename := fmt.Sprintf("excel_files/SL_%02v-%02v-%04v_%02v-%02v-%04v.xlsx", startDay, int(startMonth), startYear, endDay, int(endMonth), endYear)
	if err := f.SaveAs(filename); err != nil {
		fmt.Println(err)
		return "", err
	}

	return filename, nil
}

func coordinate(offset byte, row int) string {
	col := string([]byte("A")[0] + offset)
	return fmt.Sprintf("%v%v", col, row+1)
}

func truncateStr(str string) string {
	newStr := ""
	for i, r := range str {
		if i == 6 {
			break
		}
		newStr += fmt.Sprint(string(r))
	}
	return newStr
}

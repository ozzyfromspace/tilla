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

func generateExcel(convDoc *[]ConversionDoc, month time.Month, year int) error {
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

		// y, m, _ := (*events)[0].DateTime.Date()
		sheetName := fmt.Sprintf("%v %v - %v", student.FirstName, student.LastName, truncateStr(student.Id))
		index, err := f.NewSheet(sheetName)
		if err != nil {
			fmt.Println(err)
			return err
		}

		// create headers:
		f.SetCellValue(sheetName, "A1", "Course")
		f.SetCellValue(sheetName, "B1", "Teacher")
		f.SetCellValue(sheetName, "C1", "Date")
		f.SetCellValue(sheetName, "D1", "Time")
		f.SetCellValue(sheetName, "E1", "Duration")
		f.SetCellValue(sheetName, "F1", "Fee")
		f.SetCellValue(sheetName, "G1", "Payment Status")

		for i, event := range events {
			f.SetCellValue(sheetName, coordinate(0, i+1), event.Course)
			f.SetCellValue(sheetName, coordinate(1, i+1), event.Teacher)
			f.SetCellValue(sheetName, coordinate(2, i+1), event.Date)
			f.SetCellValue(sheetName, coordinate(3, i+1), event.Time)
			f.SetCellValue(sheetName, coordinate(4, i+1), event.Duration)
			f.SetCellValue(sheetName, coordinate(5, i+1), event.Fee)
			f.SetCellValue(sheetName, coordinate(6, i+1), event.PaymentStatus)
		}
		// Set active sheet of the workbook.
		if ci == 0 {
			f.SetActiveSheet(index)
		}
	}

	// x := (*convDoc[0]).
	// y, m, _ := (*events)[0].DateTime.Date()

	// Save spreadsheet by the given path.
	filename := fmt.Sprintf("EclipseAcademy_%02v_%04v.xlsx", month, year)
	if err := f.SaveAs(filename); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
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

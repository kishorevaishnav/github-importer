package main

import (
	"fmt"
	"log"

	"github.com/tealeg/xlsx"
)

type TestCaseStruct struct {
	Title           string
	Description     string
	TestSteps       string
	TestData        string
	ExpectedResults string
	Labels          string
	Assignee        string
}

func readXLSX(fileBytes []byte) []TestCaseStruct {
	// Open the XLSX file
	xlFile, err := xlsx.OpenBinary(fileBytes)
	if err != nil {
		log.Fatalf("Error opening XLSX file: %v", err)
	}

	var testCases []TestCaseStruct

	// Iterate through each sheet in the XLSX file
	for _, sheet := range xlFile.Sheets {
		// No need to iterate Multiple sheets as mostly its going to be only one.
		fmt.Printf("Sheet Name: %s\n", sheet.Name)

		// Iterate through each row in the sheet
		for _, row := range sheet.Rows {
			// Iterate through each cell in the row and add to the testCases array
			testCases = append(testCases, TestCaseStruct{
				Title:           row.Cells[0].Value,
				Description:     row.Cells[1].Value,
				TestSteps:       row.Cells[2].Value,
				TestData:        row.Cells[3].Value,
				ExpectedResults: row.Cells[4].Value,
				Labels:          row.Cells[5].Value,
				Assignee:        row.Cells[6].Value,
			})
		}
	}
	return testCases
}

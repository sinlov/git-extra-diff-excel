package excel_file_reader

import (
	"encoding/csv"
	"fmt"
	"github.com/gookit/color"
	"github.com/sinlov-go/unittest-kit/unittest_file_kit"
	"github.com/xuri/excelize/v2"
	"os"
)

func (e *ExcelFileReader) CheckFilePaths() error {

	if len(e.filePaths) == 0 {
		return fmt.Errorf("want read files path is empty")
	}

	for _, filePath := range e.filePaths {
		if filePath == "" {
			return fmt.Errorf("file path is empty")
		}
		if !unittest_file_kit.PathExistsFast(filePath) {
			return fmt.Errorf("file path is not excel file: %v", filePath)
		}
	}

	return nil
}

func (e *ExcelFileReader) ReadExcelFilesAsStdout() error {
	for _, filePath := range e.filePaths {
		err := e.readExcelFileAsStdout(filePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *ExcelFileReader) readExcelFileAsStdout(excelPath string) error {
	excelFile, err := excelize.OpenFile(excelPath)
	if err != nil {
		errOpenExcel := fmt.Errorf("open excel file error: %v", err)
		if e.isIgnoreRead {
			color.Redln(errOpenExcel.Error())
			return nil
		}
		return errOpenExcel
	}
	defer func() {
		// Close the spreadsheet.
		if errClose := excelFile.Close(); errClose != nil {
			fmt.Println(errClose)
		}
	}()

	for _, name := range excelFile.GetSheetMap() {
		fmt.Println("SheetName: " + name)

		WriterCSV := csv.NewWriter(os.Stdout)

		rows, errRows := excelFile.Rows(name)
		if errRows != nil {
			errReadRows := fmt.Errorf("read excel sheet error: %v", errRows)
			if e.isIgnoreRowsRead {
				color.Redln(errReadRows.Error())
				continue
			}
			return errReadRows
		}
		for rows.Next() {
			row, errColumns := rows.Columns()
			if errColumns != nil {
				errReadColumns := fmt.Errorf("read excel sheet error: %v", errColumns)
				if e.isIgnoreColumnRead {
					color.Redln(errReadColumns.Error())
					continue
				}
				return errReadColumns
			}
			errCsvWrite := WriterCSV.Write(row)
			if errCsvWrite != nil {
				errCsvWriteError := fmt.Errorf("write csv error: %v", errCsvWrite)
				if e.isIgnoreCsvWrite {
					color.Redln(errCsvWriteError.Error())
					continue
				}

				return errCsvWriteError
			}
		}
		WriterCSV.Flush()
	}

	return nil
}

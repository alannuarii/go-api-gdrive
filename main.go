package main

import (
	// "log"
	"fmt"
	// "strconv"
	// "go-api-gdrive/utils"
	// "github.com/tealeg/xlsx"
    "github.com/360EntSecGroup-Skylar/excelize"
)



func main() {

    // fileID := "1x1X-70dG8V26BZxxYBZthmdCcyu2MfoF"

    // if err := utils.DownloadFile(fileID); err != nil {
    //     log.Fatalf("Unable to download file: %v", err)
    // }

	xlsxFile := "download/07. DATA PRODUKSI HARIAN JULI 2024.xlsx"
	f, err := excelize.OpenFile(xlsxFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	sheetName := "7" // Ganti sesuai nama sheet Anda
	row := 7         // Ganti dengan nomor baris yang ingin dibaca
	col := "K"       // Ganti dengan huruf kolom yang ingin dibaca

	cellValue := f.GetCellValue(sheetName, fmt.Sprintf("%s%d", col, row))
	if cellValue == "" {
		fmt.Println("Error reading cell value:", err)
		return
	}

	fmt.Printf("Data pada sheet %s baris %d kolom %s: %s\n", sheetName, row, col, cellValue)
}


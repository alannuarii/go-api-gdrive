package main

import (
	// "log"
	"fmt"
	"time"
	"strconv"
	// "go-api-gdrive/utils"
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

	sheetName := "9" // Ganti sesuai nama sheet Anda
	row := 12         // Ganti dengan nomor baris yang ingin dibaca
	col := "H"       // Ganti dengan huruf kolom yang ingin dibaca

	cellValue := f.GetCellValue(sheetName, fmt.Sprintf("%s%d", col, row))
	if cellValue == "" {
		fmt.Println("Error reading cell value:", err)
		return
	}

	// Coba konversi nilai sel menjadi float64
	numValue, err := strconv.ParseFloat(cellValue, 64)
	if err != nil {
		// Jika tidak bisa dikonversi ke float64, mungkin nilai bukan numerik atau waktu
		fmt.Printf("Data pada sheet %s baris %d kolom %s: %s\n", sheetName, row, col, cellValue)
		return
	}

	// Periksa apakah nilai numerik ini mungkin merupakan waktu Excel
	if numValue > 0 && numValue < 1 {
		// Excel menyimpan waktu sebagai bagian desimal dari hari (misalnya, 0.5 = 12:00 PM)
		timeValue := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC).Add(time.Duration(numValue * 24 * float64(time.Hour)))
		formattedValue := timeValue.Format("15:04")
		fmt.Printf("Data pada sheet %s baris %d kolom %s: %s\n", sheetName, row, col, formattedValue)
	} else {
		// Jika tidak, format sebagai angka biasa
		formattedValue := fmt.Sprintf("%.2f", numValue)
		fmt.Printf("Data pada sheet %s baris %d kolom %s: %s\n", sheetName, row, col, formattedValue)
	}
}


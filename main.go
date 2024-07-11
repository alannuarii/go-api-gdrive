package main

import (
	// "log"
	"fmt"
	// "time"
	"strconv"
	"strings"
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

	sheetName := "Rekapitulasi" // Ganti sesuai nama sheet Anda
	row := 6                    // Ganti dengan nomor baris yang ingin dibaca
	col := "K"                  // Ganti dengan huruf kolom yang ingin dibaca

	cellValue := f.GetCellValue(sheetName, fmt.Sprintf("%s%d", col, row))
	if cellValue == "" {
		fmt.Println("Error reading cell value:", err)
		return
	}

	// Jika nilai sel dalam format HH:MM
	if strings.Contains(cellValue, ":") {
		timeParts := strings.Split(cellValue, ":")
		hours, err := strconv.Atoi(timeParts[0])
		if err != nil {
			fmt.Println("Error parsing hours:", err)
			return
		}
		minutes, err := strconv.Atoi(timeParts[1])
		if err != nil {
			fmt.Println("Error parsing minutes:", err)
			return
		}

		// Konversi ke desimal
		decimalValue := float64(hours) + float64(minutes)/60
		fmt.Printf("Data pada sheet %s baris %d kolom %s: %.2f\n", sheetName, row, col, decimalValue)
	} else {
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
			hours := int(numValue * 24)
			minutes := int((numValue*24 - float64(hours)) * 60)
			decimalValue := float64(hours) + float64(minutes)/60
			fmt.Printf("Data pada sheet %s baris %d kolom %s: %.2f\n", sheetName, row, col, decimalValue)
		} else {
			// Jika tidak, format sebagai angka biasa
			formattedValue := fmt.Sprintf("%.2f", numValue)
			fmt.Printf("Data pada sheet %s baris %d kolom %s: %s\n", sheetName, row, col, formattedValue)
		}
	}
}


package main

import (
	"log"
	"fmt"
	"strconv"
	// "go-api-gdrive/utils"
	"github.com/tealeg/xlsx"
)



func main() {

    // fileID := "1x1X-70dG8V26BZxxYBZthmdCcyu2MfoF"

    // if err := utils.DownloadFile(fileID); err != nil {
    //     log.Fatalf("Unable to download file: %v", err)
    // }

	xlFile, err := xlsx.OpenFile("download/07. DATA PRODUKSI HARIAN JULI 2024.xlsx")
    if err != nil {
        log.Fatalf("Error opening file: %v", err)
    }

    // Temukan sheet yang diinginkan
    var sheet *xlsx.Sheet
    for _, s := range xlFile.Sheets {
        if s.Name == "7" {
            sheet = s
            break
        }
    }

    if sheet == nil {
        log.Fatalf("Sheet '7' not found")
    }

	for i := 6; i < 14; i++ { // baris 7 sampai 14 (indeks dimulai dari 0)
        cell := sheet.Cell(i, 10) // Kolom K adalah kolom ke-10 (indeks dimulai dari 0)
        // Menggunakan Float64 untuk membaca angka dengan presisi yang tepat
        value, err := cell.Float()
        if err != nil {
            log.Fatalf("Error reading cell value: %v", err)
        }
        // Format angka dengan dua angka desimal
        formattedValue := strconv.FormatFloat(value, 'f', 2, 64)
        fmt.Printf("Cell K%d: %s\n", i+1, formattedValue) // i+1 karena ingin mencetak baris ke-7 sampai ke-14
    }
}

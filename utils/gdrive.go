package utils

import (
    "context"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"

    "google.golang.org/api/drive/v3"
    "google.golang.org/api/option"
)

func DownloadFile(fileID string) error {

	ctx := context.Background()

    // Path where you want to save the downloaded file
    downloadDir := "download"

    // Read the service account JSON key file
    credentialsFile := "credential/secret.json"
    srv, err := drive.NewService(ctx, option.WithCredentialsFile(credentialsFile))
    if err != nil {
        log.Fatalf("Unable to retrieve Drive client: %v", err)
    }

    // Get the metadata of the file
    fileInfo, err := srv.Files.Get(fileID).Fields("name, mimeType").Do()
    if err != nil {
        return fmt.Errorf("unable to retrieve file metadata: %v", err)
    }

    // Check if the file type is xlsx
    if fileInfo.MimeType != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
        return fmt.Errorf("file is not in XLSX format")
    }

    // Determine the file name
    fileName := fileInfo.Name
    filePath := filepath.Join(downloadDir, fileName)

    // Create a file at the specified path
    file, err := os.Create(filePath)
    if err != nil {
        return fmt.Errorf("unable to create file: %v", err)
    }
    defer file.Close()

    // Get the file content from Google Drive
    resp, err := srv.Files.Get(fileID).Download()
    if err != nil {
        return fmt.Errorf("unable to retrieve file: %v", err)
    }
    defer resp.Body.Close()

    // Write the file content to the local file
    _, err = io.Copy(file, resp.Body)
    if err != nil {
        return fmt.Errorf("unable to write file: %v", err)
    }

    fmt.Printf("File downloaded successfully to: %s\n", filePath)
    return nil
}
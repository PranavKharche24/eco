package cmd

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
)

const DownloadFolder = "/home/pranav/projects/eco/packages"

// DownloadAndRenameFile downloads from sourceURL, saving as newFileName in DownloadFolder.
func DownloadAndRenameFile(sourceURL, newFileName string) error {
    if err := os.MkdirAll(DownloadFolder, os.ModePerm); err != nil {
        return fmt.Errorf("failed to create folder %s: %v", DownloadFolder, err)
    }

    dstPath := filepath.Join(DownloadFolder, newFileName)
    fmt.Printf("Downloading from %s ...\n", sourceURL)

    resp, err := http.Get(sourceURL)
    if err != nil {
        return fmt.Errorf("failed to download file: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("download request failed with status %d", resp.StatusCode)
    }

    outFile, err := os.Create(dstPath)
    if err != nil {
        return fmt.Errorf("failed to create file %s: %v", dstPath, err)
    }
    defer outFile.Close()

    _, err = io.Copy(outFile, resp.Body)
    if err != nil {
        return fmt.Errorf("failed to write to file %s: %v", dstPath, err)
    }

    fmt.Printf("Successfully saved %s to %s\n", newFileName, DownloadFolder)
    return nil
}

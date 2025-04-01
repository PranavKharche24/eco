package cmd

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/spf13/cobra"
)

// removeCmd represents "eco remove <packageName>"
var removeCmd = &cobra.Command{
    Use:   "remove [packageName]",
    Short: "Remove a package from your system",
    Long:  "Removes the specified package from the store/packages folder (placeholder).",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        packageName := args[0]

        // In reality, you'd remove the .deb file or its contents from your local store.
        path := filepath.Join("store", "packages", packageName)

        // Example placeholder: remove file from disk (if it exists)
        err := os.Remove(path)
        if err != nil {
            fmt.Printf("Could not remove package: %s\nError: %v\n", packageName, err)
            return
        }
        fmt.Printf("Package %s removed from %s\n", packageName, path)
    },
}

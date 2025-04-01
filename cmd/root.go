package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

// rootCmd is the base command for the eco CLI
var rootCmd = &cobra.Command{
    Use:   "eco",
    Short: "eco is an advanced package manager with P2P functionality",
    Long: `eco is an advanced package manager that lets you install, update,
remove packages, and share them over a P2P connection. The actual logic for
each command is currently a placeholder, which you can fill in yourself.`,
}

// Execute runs the root command
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    // Register subcommands
    rootCmd.AddCommand(installCmd)
    rootCmd.AddCommand(updateCmd)
    rootCmd.AddCommand(removeCmd)
    rootCmd.AddCommand(shareCmd)
}

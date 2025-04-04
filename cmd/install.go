package cmd

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"
    "path/filepath"

    "github.com/spf13/cobra"
)

// PackageInfo represents a single package in the JSON
type PackageInfo struct {
    SourceURL   string `json:"source_url"`
    FileName    string `json:"file_name"`
    Description string `json:"description"`
}

// PackageDB represents the structure of our packages.json file
type PackageDB struct {
    Packages map[string]PackageInfo `json:"packages"`
}

// GetPackageJSONPath computes the path to the packages.json file
func GetPackageJSONPath() string {
    projectRoot := GetProjectRoot()
    return filepath.Join(projectRoot, "packages.json")
}

// installCmd represents "eco install <packageName>"
// For example: eco install jdk-24
var installCmd = &cobra.Command{
    Use:   "install [packageName]",
    Short: "Download and install a Debian package",
    Long: `Installs the specified package on Debian-based systems. If recognized, 
this command will first download the relevant file, rename it, and place it 
into /home/pranav/projects/eco/packages, then install it using "dpkg -i".`,
    Args: cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        packageName := args[0]

        // Load the packages from the JSON file
        packageDB, err := loadPackageDB()
        if err != nil {
            return fmt.Errorf("failed to load package database: %v", err)
        }

        // Look up the package information
        pkgInfo, ok := packageDB.Packages[packageName]
        if !ok {
            return fmt.Errorf("unrecognized package: %s", packageName)
        }

        fmt.Printf("Installing %s: %s...\n", packageName, pkgInfo.Description)

        // Download the package
        err = DownloadAndRenameFile(pkgInfo.SourceURL, pkgInfo.FileName)
        if err != nil {
            return fmt.Errorf("failed to download package %s: %v", packageName, err)
        }

        // Install the package
        err = installDebPackage(pkgInfo.FileName)
        if err != nil {
            return fmt.Errorf("failed to install package %s: %v", packageName, err)
        }

        fmt.Printf("%s installed successfully as %s.\n", packageName, pkgInfo.FileName)
        return nil
    },
}

// loadPackageDB loads the package database from the JSON file
func loadPackageDB() (*PackageDB, error) {
    data, err := ioutil.ReadFile(GetPackageJSONPath())
    if err != nil {
        return nil, fmt.Errorf("error reading package database: %v", err)
    }

    var packageDB PackageDB
    if err := json.Unmarshal(data, &packageDB); err != nil {
        return nil, fmt.Errorf("error parsing package database: %v", err)
    }

    return &packageDB, nil
}

// Add a new list command to show available packages
var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List available packages",
    RunE: func(cmd *cobra.Command, args []string) error {
        packageDB, err := loadPackageDB()
        if err != nil {
            return fmt.Errorf("failed to load package database: %v", err)
        }

        fmt.Println("Available packages:")
        fmt.Println("------------------")
        for name, pkg := range packageDB.Packages {
            fmt.Printf("- %s: %s\n", name, pkg.Description)
        }
        return nil
    },
}

// installDebPackage runs "dpkg -i" on the downloaded .deb file.
func installDebPackage(debName string) error {
    // Keep existing implementation
    debPath := filepath.Join(DownloadFolder, debName)

    fmt.Printf("Running: dpkg -i %s\n", debPath)
    cmd := exec.Command("sudo", "dpkg", "-i", debPath)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("error running dpkg -i: %v", err)
    }
    return nil
}

func init() {
    rootCmd.AddCommand(installCmd)
    rootCmd.AddCommand(listCmd) // Add the list command
}

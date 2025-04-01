package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
    "github.com/yourusername/eco/p2p"
)

var shareReceiverCmd = &cobra.Command{
    Use:   "receiver",
    Short: "Connect to a P2P server to receive packages",
    Long: `Connects to the IP displayed by "eco share sender". 
In practice, the receiver can list and download packages from the remote store.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Starting P2P receiver/client (placeholder).")
        p2p.StartClient()
    },
}

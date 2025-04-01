package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
    "github.com/yourusername/eco/p2p"
)

var shareSenderCmd = &cobra.Command{
    Use:   "sender",
    Short: "Start a P2P server to share packages",
    Long: `Starts a server that displays its IP. Another peer can run "eco share receiver"
to connect and see or download packages from the store/packages folder.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Starting P2P sender/server (placeholder).")
        p2p.StartServer()
    },
}

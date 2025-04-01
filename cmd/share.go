package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var shareCmd = &cobra.Command{
    Use:   "share",
    Short: "Manage P2P sharing of packages",
    Long: `Use "eco share sender" to start a server that displays an IP on the terminal,
allowing a peer to connect. Use "eco share receiver" to connect to that IP
and fetch or list available packages.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println(`Use "eco share sender" or "eco share receiver"`)
    },
}

func init() {
    shareCmd.AddCommand(shareSenderCmd)
    shareCmd.AddCommand(shareReceiverCmd)
}

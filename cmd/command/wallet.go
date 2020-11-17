package command

import (
	"github.com/spf13/cobra"
	"github.com/zhubby/blockchain-go"
)

var (
	walletNodeId string

	walletCmd = &cobra.Command{
		Use: "wallet",
	}

	walletNewCmd = &cobra.Command{
		Use: "new",
		Run: func(cmd *cobra.Command, args []string) {
			blockchain.NewWallet()
		},
	}
)

func init() {
	walletNewCmd.PersistentFlags().StringVar(&walletNodeId, "node-id", "", "")

	walletCmd.AddCommand(walletNewCmd)

}

package cmd

import (
	"context"
	"log"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(expireCronCommand)
}

var expireCronCommand = &cobra.Command{
	Use:   "expire-payment-codes",
	Short: "Command to start job to expire payment codes",
	Run:   expirePaymentCodes,
}

func expirePaymentCodes(cmd *cobra.Command, args []string) {
	for i := 1; i < 12; i++ {
		err := expirePaymentCodesJob.Work(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)
	}
}

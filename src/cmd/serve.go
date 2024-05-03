package cmd

import (
	"github.com/s7010390/testApi/logger"
	"github.com/spf13/cobra"
)

var EnquiryCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start csv-thinker-analyst Service",
	Run: func(cmd *cobra.Command, args []string) {

		// Init Config
		initConfig()

		// Init Log
		logger.InitLogger()

		initListenOsSignal()

		// Init Interface
		initListenInterface()

		wg.Wait()

		logger.SyncLogger()
	},
}

func init() {
	rootCmd.AddCommand(EnquiryCmd)
}

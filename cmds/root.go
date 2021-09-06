package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var verbose bool

var rootCommand = &cobra.Command{
	Use:   "cmdb",
	Short: "cmdb program",
	Long:  "cmdb program",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("cmdb")
		return nil
	},
}

func init() {
	rootCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", true, "verbose")
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}

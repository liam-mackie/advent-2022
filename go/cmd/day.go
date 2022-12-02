package cmd

import (
	"errors"
	"github.com/liam-mackie/advent2022/pkg/advent"
	"github.com/liam-mackie/advent2022/pkg/days/one"

	"github.com/spf13/cobra"
)

// dayCmd represents the day command
var dayCmd = &cobra.Command{
	Use:   "day",
	Short: "Run a day's command",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a day to run")
		}
		err := advent.IsValidDay(args[0])
		if err == nil {
			return nil
		}
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "one":
			one.Run()
		}
	},
}

func init() {
	rootCmd.AddCommand(dayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

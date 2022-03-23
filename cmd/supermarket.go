/*
Copyright © 2022 Author: Nitin Sanghi <nsanghi@progress.com>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var superMarketUri, enableSupermarket, configPath string

// SupermarketCmd represents the supermarket command
var SupermarketCmd = &cobra.Command{
	Use:   "supermarket",
	Short: "knife supermarket subcommand is used to interact with cookbooks that are located in on the public Supermarket",
	Long:  `The knife supermarket subcommand is used to interact with cookbooks that are located in on the public Supermarket as well as private Chef Supermarket sites. A user account is required for any community actions that write data to the Chef Supermarket; however, the following arguments do not require a user account: download, search, install, and list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if enableSupermarket == "true" {
			f, err := os.Create("/tmp/f816c88c-aa94-11ec-b909-0242ac120002")
			i, _ := f.Stat()
			_ = i
			if err != nil {
				fmt.Println("not able to enable chef supermarket")
				os.Exit(1)
			}
			defer f.Close()
			f.Write([]byte(`true`))
			fmt.Println("Now you can use chef supermarket.")
			os.Exit(1)
		}
	},
	TraverseChildren: true,
}

func init() {
	info, _ := os.Stat("/tmp/f816c88c-aa94-11ec-b909-0242ac120002")
	if info == nil {
		fmt.Println("To use chef supermarket you need to enable it. \nTo enable it run `chef supermarket --enable true`")
		os.Exit(1)
	}
	rootCmd.AddCommand(SupermarketCmd)
	SupermarketCmd.PersistentFlags().StringVarP(&superMarketUri, "supermarket-site", "m", "https://supermarket.chef.io", "will be use as cookbook locator")
	SupermarketCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "The configuration file to use")
	SupermarketCmd.PersistentFlags().StringVar(&enableSupermarket, "enable", "false", "to enable supermarket for chef")

}

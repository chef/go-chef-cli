package cmd

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/chef/go-knife/core"
	"github.com/chef/go-knife/dsl"
	"github.com/chef/go-knife/supermarket"
	"github.com/spf13/cobra"
)

var (
	noDeps           bool
	cookBookPath     string
	defaultBranch    string
	useCurrentBranch bool
)

// supermarketSearchCmd represents the supermarket search
var supermarketInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Search indexes allow queries to be made for any type of data that is indexed by the Chef Infra Server, including data bags (and data bag items), environments, nodes, and roles",
	Long:  `Search indexes allow queries to be made for any type of data that is indexed by the Chef Infra Server, including data bags (and data bag items), environments, nodes, and roles. A defined query syntax is used to support search patterns like exact, wildcard, range, and fuzzy. A search is a full-text query that can be done from several locations, including from within a recipe, by using the search subcommand in knife.`,
	Run: func(cmd *cobra.Command, args []string) {
		var ui core.UI
		supermarket.ValidateArgsAndType(args, "", ui)
		if !supermarket.ValidateArtifact(args[0]) {
			ui.Msg("only cookbook type artifact supported as of now.")
			os.Exit(1)
		}
		if len(args) >= 3 {
			version, _ := regexp.MatchString(`(\d+)(\.\d+){1,2}`, args[2])
			if !version {
				ui.Fatal("Installing multiple cookbooks at once is not supported.")
			}
		}
		installPath := ""
		if len(cookBookPath) > 1 {
			cookbooks := strings.Split(cookBookPath, ":")
			installPath = cookbooks[0]
		} else {
			installPath = core.GetDefaultConfigPath()
		}
		var config core.Config
		config.Format = format
		ci := supermarket.NewInstallProvider(args[1], superMarkerUri, installPath, defaultBranch, args[0], noDeps, useCurrentBranch)
		ci.Install(ui, config)
		if !ci.InstallDeps() {
			m, err := dsl.ReadMetaData(filepath.Join(installPath, args[1]))
			if err != nil {
				ui.Error("unable to read meta file: " + err.Error())
				os.Exit(1)
			}
			for name := range m.Dependencies {
				ci.ChangeArtifactName(name)
				ci.Install(ui, config)
			}
		}
	},
}

func init() {
	SupermarketCmd.AddCommand(supermarketInstallCmd)
	// supermarketInstallCmd.PersistentFlags().StringVarP(&superMarkerUri, "supermarket-site", "m", "https://supermarket.chef.io", "will be use to search cookbook")
	supermarketInstallCmd.PersistentFlags().StringVarP(&cookBookPath, "cookbook-path", "o", "", "A colon-separated path to look for cookbooks in.")
	supermarketInstallCmd.PersistentFlags().StringVarP(&defaultBranch, "branch", "B", "master", "Default branch to work with.")
	supermarketInstallCmd.PersistentFlags().BoolVarP(&noDeps, "skip-dependencies", "D", false, "Skips automatic dependency installation.")
	supermarketInstallCmd.PersistentFlags().BoolVarP(&useCurrentBranch, "use-current-branch", "b", false, "Use the current branch.")

}

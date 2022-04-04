package supermarket

import (
	"os"
	"testing"

	"github.com/chef/go-chef-cli/core"
	"github.com/go-chef/chef"
)

func TestCookbookUnshare(t *testing.T) {

	var config core.Config
	var ui core.UI

	homeDir, _ := os.UserHomeDir()
	data, path := config.LoadConfig(homeDir + "/.chef/")
	var cc chef.ConfigRb
	cc, err := chef.NewClientRb(data, path)
	if err != nil {
		ui.Fatal("No chef configuration file found. See https://docs.chef.io/config_rb/ for details.")
	}
	err = UnShareCookbook("mysql", "https://supermarket.chef.io", cc.NodeName, cc.ClientKey)

	if err == nil {
		t.Error("not valid output for  mysql")
	}

}

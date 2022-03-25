package supermarket

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/chef/go-chef-cli/core"
)

func TestCookbookShow(t *testing.T) {

	var config core.Config
	var ui core.UI

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	ShowCookBook("mysql", "https://supermarket.chef.io", ui, config, "8.5.1")
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	data := string(out)
	if !strings.Contains(data, "https://supermarket.chef.io/api/v1/cookbooks/mysql") {
		t.Error("not valid output for  mysql")
	}
	if !strings.Contains(data, "https://supermarket.chef.io/api/v1/cookbooks/mysql") {
		t.Error("not valid output for  mysql")
	}
	if !strings.Contains(data, "2017-08-23T19:01:28Z") {
		t.Error("not valid output for mysql")
	}

}

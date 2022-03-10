package supermarket

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/chef/go-chef-cli/core"
)

func TestCookbookSearch_Search(t *testing.T) {
	var cs CookbookSearch
	cs.Start = 0
	cs.Rows = searchItemsCount
	cs.Query = "mysql"
	cs.Url = "https://supermarket.chef.io"
	expectedResult := make(map[string]cookbook, 10)
	expectedResult["L7-mysql"] = cookbook{
		Name:       "L7-mysql",
		Maintainer: "szelcsanyi",
		Cookbook:   "https://supermarket.chef.io/api/v1/cookbooks/l7-mysql",
	}
	expectedResult["cg_mysql"] = cookbook{
		Name:       "cg_mysql",
		Maintainer: "phai",
		Cookbook:   "https://supermarket.chef.io/api/v1/cookbooks/cg_mysql",
	}
	expectedResult["mw_mysql"] = cookbook{
		Name:       "mw_mysql",
		Maintainer: "car",
		Cookbook:   "https://supermarket.chef.io/api/v1/cookbooks/mw_mysql",
	}
	expectedResult["mysql"] = cookbook{
		Name:       "mysql",
		Maintainer: "sous-chefs",
		Cookbook:   "https://supermarket.chef.io/api/v1/cookbooks/mysql",
	}
	expectedResult["mysql-mha"] = cookbook{
		Name:       "mysql-mha",
		Maintainer: "ovaistariq",
		Cookbook:   "https://supermarket.chef.io/api/v1/cookbooks/mysql-mha",
	}
	expectedResult["mysql-multi"] = cookbook{
		Name:       "mysql-multi",
		Maintainer: "rackops",
		Cookbook:   "https://supermarket.chef.io/api/v1/cookbooks/mysql-multi",
	}
	expectedResult["mysql-sys"] = cookbook{
		Name:       "mysql-sys",
		Maintainer: "ovaistariq",
		Cookbook:   "https://supermarket.chef.io/api/v1/cookbooks/mysql-sys",
	}
	var config core.Config
	var ui core.UI

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	cs.Search(ui, config)
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	result := make(map[string]cookbook, 10)
	err := json.Unmarshal(out, &result)
	if err != nil {
		t.Error("not valid output for search query mysql")
	}
	for name, c := range result {
		if data, ok := expectedResult[name]; ok && !checkCookbookData(data, c) {
			out := fmt.Sprintf("expected: %v, got: %v", data, c)
			t.Error("invalid search result output does not added up", out)
		}
	}
}

func TestCookbookSearch_String(t *testing.T) {
	var cs CookbookSearch
	cs.Start = 0
	cs.Rows = searchItemsCount
	cs.Query = "mysql"
	cs.Url = "https://supermarket.chef.io"
	if cs.String() != "https://supermarket.chef.io/api/v1/search?q=mysql&rows=1000&start=0" {
		t.Error("invalid cookbook search uri")
	}

}

func TestCookbookNoData(t *testing.T) {
	var cs CookbookSearch
	cs.Start = 0
	cs.Rows = searchItemsCount
	cs.Query = ""
	cs.Url = "https://supermarket.chef.io"
	var config core.Config
	var ui core.UI

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	cs.Search(ui, config)
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	result := make(map[string]cookbook, 10)
	err := json.Unmarshal(out, &result)
	if err != nil {
		t.Error("not valid output for search query mysql")
	}
	if len(result) > 0 {
		t.Error("result found for no query")
	}
}

func checkCookbookData(a, b cookbook) bool {
	if a.Maintainer != b.Maintainer {
		return false
	}
	if a.Name != b.Name {
		return false
	}
	if a.Cookbook != b.Cookbook {
		return false
	}
	return true
}

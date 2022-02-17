package supermarket

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/chef/go-knife/core"
)

var ui core.UI
var config core.Config

func TestCookbookDownload_Deprecated(t *testing.T) {
	var cd CookbookDownload

	cd.da.Url = "https://supermarket.chef.io"
	cd.da.ArtifactName = "capistrano"
	cd.CookbookName = "capistrano"
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	err := cd.Download(ui, config)
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if err == nil {
		t.Error("downloading deprecated cookbook it should give WARNING")
	}
	if !strings.Contains(string(out), "It has been replaced by") {
		t.Error("not showing replaced cookbook")
	}
	cd.da.Force = true
	err = cd.Download(ui, config)
	if err != nil {
		t.Error("cookbook should download with force")
	}
}
func TestCookbookDownload_DownloadWithForce(t *testing.T) {
	var cd CookbookDownload

	cd.da.Url = "https://supermarket.chef.io"
	cd.da.ArtifactName = "capistrano"
	cd.CookbookName = "capistrano"
	cd.da.Force = true
	err := cd.Download(ui, config)
	if err != nil {
		t.Error("cookbook should download with force")
	}
	cookbookPath := filepath.Join("./", "capistrano-0.7.0.tar.gz")
	info, err := os.Stat(cookbookPath)
	if err != nil && info.Name() == "capistrano-0.7.0.tar.gz" {
		t.Error("not able to save deprecated cookbook ")

	}
	os.Remove(cookbookPath)

}
func TestCookbookDownload_Download(t *testing.T) {
	var cd CookbookDownload

	cd.da.Url = "https://supermarket.chef.io"
	cd.da.ArtifactName = "mysql"
	cd.CookbookName = "mysql"
	err := cd.Download(ui, config)
	if err != nil {
		t.Error("cookbook should download")
	}
	cookbookPath := filepath.Join("./", "mysql-11.0.4.tar.gz")
	info, err := os.Stat(cookbookPath)
	if err != nil && info.Name() == "mysql-11.0.4.tar.gz" {
		t.Error("not able to save deprecated cookbook ")

	}
	os.Remove(cookbookPath)

}
func TestCookbookDownload_Version(t *testing.T) {
	var cd CookbookDownload
	cd.version = "0.0.0"
	if cd.Version() != "0.0.0" {
		t.Error("invalid cookbook version")
	}
}
func TestCookbookDownload_String(t *testing.T) {
	var cd CookbookDownload

	cd.da.Url = "https://supermarket.chef.io"
	cd.da.ArtifactName = "mysql"
	if cd.String() != "https://supermarket.chef.io/api/v1/cookbooks/mysql" {
		t.Error("invalid cookbook download uri")
	}
}

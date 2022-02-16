package dsl

import (
	"testing"
)

const _IssueUrl = "https://github.com/<insert_org_here>/apache/issues"
const _Name = "apache"
const _Maintainer = "The Authors"
const _MaintainerEmail = "you@example.com"
const _SourceUrl = "https://github.com/<insert_org_here>/apache"
const _License = "All Rights Reserved"
const _Version = "0.1.0"
const _ChefVersion = ">= 15.0"
const _Description = "Installs/Configures apache"

func TestNewMetaData(t *testing.T) {
	data := `   name 'apache'
				maintainer 'The Authors'
				maintainer_email 'you@example.com'
				license 'All Rights Reserved'
				description 'Installs/Configures apache'
				version '0.1.0'
				chef_version '>= 15.0'
				
				#issues_url points to the location where issues for this cookbook are
				# tracked.  A View Issues link will be displayed on this cookbook's page when
				# uploaded to a Supermarket.
				#
				issues_url 'https://github.com/<insert_org_here>/apache/issues'
				
				source_url 'https://github.com/<insert_org_here>/apache'`
	md, err := NewMetaData(data)
	if err != nil {
		t.Error("invalid metadata.rb contain please validate it", err)
	}
	validateMetaData(md, t, "TestNewMetaData")

}

func TestNewMetaDataFromJson(t *testing.T) {
	data := `{"name":"apache","description":"Installs/Configures apache","long_description":"","maintainer":"The Authors","maintainer_email":"you@example.com","license":"All Rights Reserved","platforms":{},"dependencies":{},"providing":null,"recipes":null,"version":"0.1.0","source_url":"https://github.com/\u003cinsert_org_here\u003e/apache","issues_url":"https://github.com/\u003cinsert_org_here\u003e/apache/issues","ChefVersion":"\u003e= 15.0","OhaiVersion":"","gems":null,"eager_load_libraries":false,"privacy":false}`
	md, err := NewMetaDataFromJson([]byte(data))
	if err != nil {
		t.Error("invalid metadata.rb contain please validate it", err)
	}
	validateMetaData(md, t, "TestNewMetaDataFromJson")
}
func TestReadMetaData(t *testing.T) {
	// file, err := os.Create("/tmp/metadata.rb")
	// if err != nil {
	// 	t.Error("unable to create tmo metadata.rb", err)
	// }
}
func TestReadMetaData2(t *testing.T) {

}
func validateMetaData(md MetaData, t *testing.T, funcName string) {
	if md.Description != _Description {
		t.Errorf("%s: invalid Description", funcName)
	}
	if md.IssueUrl != _IssueUrl {
		t.Errorf("%s: invalid IssueUrl", funcName)
	}
	if md.Name != _Name {
		t.Errorf("%s: invalid Name", funcName)
	}
	if md.Maintainer != _Maintainer {
		t.Errorf("%s: invalid Maintainer", funcName)
	}
	if md.MaintainerEmail != _MaintainerEmail {
		t.Errorf("%s: invalid MaintainerEmail", funcName)
	}
	if md.SourceUrl != _SourceUrl {
		t.Errorf("%s: invalid SourceUrl", funcName)
	}
	if md.License != _License {
		t.Errorf("%s: invalid License", funcName)
	}
	if md.Version != _Version {
		t.Errorf("%s: invalid Version", funcName)
	}
	if md.ChefVersion != _ChefVersion {
		t.Errorf("%s: invalid ChefVersion", funcName)
	}

}

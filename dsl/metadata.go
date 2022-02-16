package dsl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const metaRbName = "metadata.rb"
const metaJsonName = "metadata.json"

type metaFunc func(s []string, m *MetaData)

var metaRegistry map[string]metaFunc

func init() {
	metaRegistry = make(map[string]metaFunc, 15)
	metaRegistry["name"] = metaNameParser
	metaRegistry["maintainer"] = metaMaintainerParser
	metaRegistry["maintainer_email"] = metaMaintainerMailParser
	metaRegistry["license"] = metaLicenseParser
	metaRegistry["description"] = metaDescriptionParser
	metaRegistry["long_description"] = metaLongDescriptionParser
	metaRegistry["source_url"] = metaSourceUrlParser
	metaRegistry["issues_url"] = metaIssueUrlParser
	metaRegistry["platforms"] = metaSupportsParser
	metaRegistry["supports"] = metaSupportsParser
	metaRegistry["%w("] = metaSupportsRubyParser
	metaRegistry["privacy"] = metaPrivacyParser
	metaRegistry["depends"] = metaDependsParser
	metaRegistry["version"] = metaVersionParser
	metaRegistry["chef_version"] = metaChefVersionParser
	metaRegistry["ohai_version"] = metaOhaiVersionParser
	metaRegistry["gem"] = metaGemParser

}

// MetaData metadata mapping of metadata.rb or metadata.json
type MetaData struct {
	Name               string            `json:"name"`
	Description        string            `json:"description"`
	LongDescription    string            `json:"long_description"`
	Maintainer         string            `json:"maintainer"`
	MaintainerEmail    string            `json:"maintainer_email"`
	License            string            `json:"license"`
	Platforms          map[string]string `json:"platforms"`
	Dependencies       map[string]string `json:"dependencies"`
	Providing          map[string]string `json:"providing"`
	Recipes            map[string]string `json:"recipes"`
	Version            string            `json:"version"`
	SourceUrl          string            `json:"source_url"`
	IssueUrl           string            `json:"issues_url"`
	ChefVersion        string
	OhaiVersion        string
	Gems               []string `json:"gems"`
	EagerLoadLibraries bool     `json:"eager_load_libraries"`
	Privacy            bool     `json:"privacy"`
}

func ReadMetaData(path string) (m MetaData, err error) {
	fileName := filepath.Join(path, metaJsonName)
	jsonType := true
	if !isFileExists(fileName) {
		jsonType = false
		fileName = filepath.Join(path, metaRbName)

	}
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if jsonType {
		return NewMetaDataFromJson(file)
	} else {
		return NewMetaData(string(file))
	}

}

func NewMetaData(data string) (m MetaData, err error) {
	linesData := strings.Split(data, "\n")
	if len(linesData) < 3 {
		return m, errors.New("not much info")
	}
	m.Dependencies = make(map[string]string, 1)
	m.Platforms = make(map[string]string, 1)
	for _, i := range linesData {
		key, value := getKeyValue(strings.TrimSpace(i))
		if fn, ok := metaRegistry[key]; ok {
			fn(value, &m)
		}
	}
	return m, err
}

func NewMetaDataFromJson(data []byte) (m MetaData, err error) {
	err = json.Unmarshal(data, &m)
	return m, err
}

func StringParserForMeta(s []string) string {
	str := strings.Join(s, " ")
	return trimQuotes(strings.TrimSpace(str))
}
func metaNameParser(s []string, m *MetaData) {
	m.Name = StringParserForMeta(s)
}

func metaMaintainerParser(s []string, m *MetaData) {
	m.Maintainer = StringParserForMeta(s)
}
func metaMaintainerMailParser(s []string, m *MetaData) {
	m.MaintainerEmail = StringParserForMeta(s)
}
func metaLicenseParser(s []string, m *MetaData) {
	m.License = StringParserForMeta(s)
}
func metaDescriptionParser(s []string, m *MetaData) {
	m.Description = StringParserForMeta(s)
}
func metaLongDescriptionParser(s []string, m *MetaData) {
	m.LongDescription = StringParserForMeta(s)
}
func metaIssueUrlParser(s []string, m *MetaData) {
	m.IssueUrl = StringParserForMeta(s)
}
func metaSourceUrlParser(s []string, m *MetaData) {
	m.SourceUrl = StringParserForMeta(s)
}
func metaGemParser(s []string, m *MetaData) {
	m.Gems = append(m.Gems, StringParserForMeta(s))
}

func metaVersionParser(s []string, m *MetaData) {
	m.Version = StringParserForMeta(s)
}
func metaOhaiVersionParser(s []string, m *MetaData) {
	m.OhaiVersion = StringParserForMeta(s)
}
func metaChefVersionParser(s []string, m *MetaData) {
	m.ChefVersion = StringParserForMeta(s)
}
func metaPrivacyParser(s []string, m *MetaData) {
	if s[0] == "true" {
		m.Privacy = true
	}
}
func metaSupportsParser(s []string, m *MetaData) {
	s = clearWhiteSpace(s)
	switch len(s) {
	case 1:
		if s[0] != "os" {
			m.Platforms[strings.TrimSpace(s[0])] = ">= 0.0.0"
		}
	case 2:
		m.Platforms[strings.TrimSpace(s[0])] = s[1]
	case 3:
		v := trimQuotes(s[1] + " " + s[2])
		m.Platforms[strings.TrimSpace(s[0])] = v

	}
	if len(s) > 3 {
		panic(`<<~OBSOLETED
		The dependency specification syntax you are using is no longer valid. You may not
		specify more than one version constraint for a particular cookbook.
			Consult https://docs.chef.io/config_rb_metadata/ for the updated syntax.`)
	}
}
func metaDependsParser(s []string, m *MetaData) {
	s = clearWhiteSpace(s)
	switch len(s) {
	case 1:
		m.Dependencies[strings.TrimSpace(s[0])] = ">= 0.0.0"
	case 2:
		m.Dependencies[strings.TrimSpace(s[0])] = s[1]

	case 3:
		v := trimQuotes(s[1] + " " + s[2])
		m.Dependencies[strings.TrimSpace(s[0])] = v

	}
	if len(s) > 3 {
		panic(`<<~OBSOLETED
		The dependency specification syntax you are using is no longer valid. You may not
		specify more than one version constraint for a particular cookbook.
			Consult https://docs.chef.io/config_rb_metadata/ for the updated syntax.`)
	}
}

func metaSupportsRubyParser(s []string, m *MetaData) {
	if len(s) > 1 {
		for _, i := range s {
			switch i {
			case ").each":
				continue
			case "do":
				continue
			case "|os|":
				continue
			default:
				m.Platforms[strings.TrimSpace(s[0])] = ">= 0.0.0"
			}
		}
	}
}

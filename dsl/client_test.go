package dsl

import (
	"strings"
	"testing"
)

func TestNewClientRb(t *testing.T) {
	clientRegistry["client_key"] = func(s []string, path string, m *ClientRb) {
		str := StringParserForMeta(s)
		data := strings.Split(str, "/")
		m.ClientKey = data[len(data)-1]
	}
	data := `current_dir = File.dirname(__FILE__)
		log_level                :info
		log_location             STDOUT
		node_name                "nitin"
		client_key               "#{current_dir}/nitin.pem"
		chef_server_url          "https://server/organizations/sanghi"
		cookbook_path            ["#{current_dir}/../cookbooks"]`
	cb, err := NewClientRb(data, "")
	if err != nil {
		t.Error("unable to read config.rb file")
	}
	if cb.ClientKey != "nitin.pem" {
		t.Error("client.pem have invalid path")
	}
	if cb.ChefServerUrl != "https://server/organizations/sanghi" {
		t.Error("invalid chef server url read from config.rb")
	}
}

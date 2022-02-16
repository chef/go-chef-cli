package dsl

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type ClientRb struct {
	ClientKey     string
	ChefServerUrl string
}

type clientFunc func(s []string, path string, m *ClientRb)

var clientRegistry map[string]clientFunc

func init() {
	clientRegistry = make(map[string]clientFunc, 2)
	clientRegistry["client_key"] = clientKeyParser
	clientRegistry["chef_server_url"] = clientServerParser

}
func NewClientRb(data, path string) (c ClientRb, err error) {
	linesData := strings.Split(data, "\n")
	if len(linesData) < 3 {
		return c, errors.New("not much info")
	}
	for _, i := range linesData {
		key, value := getKeyValue(strings.TrimSpace(i))
		if fn, ok := clientRegistry[key]; ok {
			fn(value, path, &c)
		}
	}
	return c, err
}
func clientKeyParser(s []string, path string, c *ClientRb) {
	str := StringParserForMeta(s)
	data := strings.Split(str, "/")
	size := len(data)
	if size > 0 {
		keyPath := filepath.Join(path, data[size-1])
		keyData, err := ioutil.ReadFile(keyPath)
		if err != nil {
			fmt.Println("error in reading pem file at: ", keyPath)
			os.Exit(1)
		}
		c.ClientKey = string(keyData)
	}
}
func clientServerParser(s []string, path string, c *ClientRb) {
	c.ChefServerUrl = StringParserForMeta(s)
}

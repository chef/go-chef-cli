package cmd

import (
	"testing"
)

func Test(t *testing.T) {
	err := rootCmd.Execute()
	if err != nil {
		t.Error("some issue the root cmd")
	}
}

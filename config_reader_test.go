package main

import "testing"

func TestReadConfig(t *testing.T) {
	fs := ReadConfig()

	if len(fs.filePattern) > 0 {
		t.Errorf("No filePattern given in config.")
	}

	if len(fs.i18nPath) > 0 {
		t.Errorf("No i18nPath given in config.")
	}
}

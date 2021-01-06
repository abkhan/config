package test

import (
	"flag"
	"os"
	"testing"

	"github.com/abkhan/config"
)

func TestValidateConfig(t *testing.T) {
	os.Args = []string{os.Args[0], "-confile=config.yml"}
	flag.Parse()
	validateConfig(t)
}

func validateConfig(t *testing.T) {
	if err := config.Load(&conf); err != nil {
		t.Errorf("Failed to load config values, %v", err)
	}
	if err := config.ValidateConf(conf); err != nil {
		t.Error(err)
	}
}

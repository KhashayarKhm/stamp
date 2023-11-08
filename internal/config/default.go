package config

import (
	"os"
	"path/filepath"
)

func Default() *Config {
	var stampDir string
	if homeDir, err := os.UserHomeDir(); err != nil {
		panic(err)
	} else {
		stampDir = filepath.Join(homeDir, ".stamp")
	}

	return &Config{
		WatermarkImg: filepath.Join(stampDir, "default.png"),
		StampDir:  stampDir,
	}
}

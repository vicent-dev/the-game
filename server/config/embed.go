package config

import (
	"embed"
)

//go:embed config.yaml
var f embed.FS

const (
	configFileName = "config.yaml"
)

func getConfigFile() []byte {
	bs, _ := f.ReadFile(configFileName)
	return bs
}

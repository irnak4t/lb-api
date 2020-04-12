package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Db DbConfig `toml:"mysql"`
}

type DbConfig struct {
	Database string `toml:"database"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

var config Config

func LoadToml() {
	exe, _ := os.Executable()

	dir := filepath.Dir(exe)
	_, _ = toml.DecodeFile(dir+"/db/db.toml", &config)
}

func Get() Config {
	LoadToml()
	return config
}

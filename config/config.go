package config

import (
	"fmt"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Bar string
	Baz string
}

func Foo() {
	c := Config{}
	_, err := toml.DecodeFile(filepath.Join("config", "default.toml"), &c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}
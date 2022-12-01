package main

import (
	"github.com/farzanehshahi/user-kit/internal/app"
)

const (
	configPath = "./config/local.yml"
)

func main() {
	app.Run(configPath)
}

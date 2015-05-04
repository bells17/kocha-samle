package main

import (
	"kocha-sample/config"

	"github.com/naoina/kocha"
)

func main() {
	if err := kocha.Run(config.AppConfig); err != nil {
		panic(err)
	}
}

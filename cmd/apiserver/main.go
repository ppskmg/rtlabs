package main

import (
	"log"
	"rtlabs/internal/app/apiserver"
)

func main() {
	config := apiserver.NewConfig()
	if err := apiserver.Start(config, true); err != nil {
		log.Fatal(err)
	}
}

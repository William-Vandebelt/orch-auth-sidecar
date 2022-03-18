package main

import (
	"log"

	_ "github.com/quasilyte/go-ruleguard/dsl"
	kit_utils "github.com/sailsforce/gomicro-kit/utils"
)

func init() {
	if err := kit_utils.InitEnv(); err != nil {
		log.Printf("error loading .env file: %v\n", err)
	}
}

func main() {
	log.Println("hello world")
}

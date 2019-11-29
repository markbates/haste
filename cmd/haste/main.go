package main

import (
	"fmt"
	"log"

	"github.com/markbates/haste"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	h, err := haste.New(".")
	if err != nil {
		return err
	}

	for _, f := range h.Funcs() {
		fmt.Println(f)
	}
	return nil
}

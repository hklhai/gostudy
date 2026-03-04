package main

import (
	"log"
	"os"
)

func main() {
	name := "test"
	f, err := os.Open(name)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if d, err := f.Stat(); err != nil {
		f.Close()
		_ = d
		log.Println(err)
		os.Exit(1)
	}
	log.Println(err)

}

package main

import "log"

func main() {
	err := Start()
	if err != nil {
		log.Fatal(err.Error())
	}
}

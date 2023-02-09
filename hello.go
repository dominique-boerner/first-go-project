package main

import "log"

func main() {
	config := ReadConfig()
	log.Printf("%#v", config)
}

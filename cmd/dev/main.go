package main

import "log"

func main() {
	if err := RunDev(); err != nil {
		log.Fatal(err)
	}
}
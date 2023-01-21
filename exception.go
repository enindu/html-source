package main

import "log"

func handle(exception error) {
	// Exit if error throws
	if exception != nil {
		log.Fatal(exception)
	}
}

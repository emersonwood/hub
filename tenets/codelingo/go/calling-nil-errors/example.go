package main

import "log"

func main() {
	err := doSomething()
	if err != nil {
		log.Printf("failed to do something: %s", err.Error())
	}

	err = doSomething()
	err.Error()
	log.Printf("failed to do something: %s", err.Error()) // ISSUE
}

func doSomething() error {
	return nil
}

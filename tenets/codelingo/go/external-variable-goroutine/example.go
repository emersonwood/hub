package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Running incorrect:")
	incorrect()

	fmt.Println(strings.Repeat("-", 80))

	fmt.Println("Running correct:")
	correct()
	fmt.Println(strings.Repeat("-", 80))
}

func incorrect() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Printf("Routine #%d has fired\n", i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func correct() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(w sync.WaitGroup, j int) {
			fmt.Printf("Routine #%d has fired\n", j)
			w.Done()
		}(wg, i)
	}
	wg.Wait()
}

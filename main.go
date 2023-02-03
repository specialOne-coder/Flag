package main

import (
	//"net/http"
	"fmt"
	"sync"
	//"net/http"
	// "sync"
	// "net/http"
	// "sync"
	// "time"
)

func main() {
	fmt.Println("Start")
	fmt.Println("")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ThirdApi()
	}()
	wg.Wait()
	fmt.Println("End API")
	fmt.Println("")
}

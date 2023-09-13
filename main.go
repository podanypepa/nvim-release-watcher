// Package main _
package main

import (
	"fmt"
	"log"
	"sync"
)


func main() {
	version, createdTime := "", ""

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		var err error
		version, err = getVersion()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		createdTime, err = getReleaseFileTime()
		if err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()

	fmt.Println("version:    ", version)
	fmt.Println("createdTime:", createdTime)
}

package function

import (
	"fmt"
	"sync"

	"github.com/skywalkeretw/master-api/app/utils"
)

func CreateFunction() {
	// Create a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	tempDirPath, err := utils.GenerateTempFolder()
	if err != nil {
		// handle error
	}
	fmt.Println(tempDirPath)
	// Increment the WaitGroup counter for each goroutine
	wg.Add(2)

	// Create AsyncAPI file
	go func() {
		defer wg.Done()
		//firstFunction()
	}()

	// Create AsyncAPI file
	go func() {
		defer wg.Done()
		//secondFunction()
	}()

	// Wait for both goroutines to finish
	wg.Wait()

	// Continue with the next command or operation
	fmt.Println("Both functions have completed.")
}

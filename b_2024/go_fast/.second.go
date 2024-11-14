package main

import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var wg = sync.WaitGroup{}
var results = []string{}

func main() {
	fmt.Println("Second part - Goroutines and forward")
	/*https://youtu.be/8uiZC0l4Ajw?si=CavnMWrtfKLWAVwY&t=2407
	- POTENTIAL NEW SOURCE OF INFORMATION
		https://youtu.be/N2GWXuj_IWg?si=bLJl86kXxqZwBPmZ&t=8966
	*/

	// Go routines - Launch multiple functions and run them concurrently
	// concurrency != parallel execution (not only)
	// it means "multiple tasks in progress at the same time"

	// one cpu core
	// i.e 2 tasks, one has a more demanding subtasks (connecting to database, for example)
	// while the task 1 (db connect) moves on, task 2 can continue its execution

	// another example -- parallel execution ( +1 cpu core)
	// simultaneously executed
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nResults are: %v", results)
}

func dbCall(i int) {
	// Simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is ", dbData[i])
	results = append(results, dbData[i])
	wg.Done()
}

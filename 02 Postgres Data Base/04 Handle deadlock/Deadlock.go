package main

import (
	"fmt"
	"sync"
	"time"
)

type Resource struct {
	ID   int
	Lock sync.Mutex
}

func main() {
	// Create resources
	resource1 := &Resource{ID: 1}
	resource2 := &Resource{ID: 2}

	go performOperations(resource1, resource2)
	go performOperations(resource2, resource1)

	time.Sleep(5 * time.Second)
}

func performOperations(r1, r2 *Resource) {
	for {
		//resource 1
		r1.Lock.Lock()
		fmt.Printf("Resource %d acquired lock\n", r1.ID)

		time.Sleep(1 * time.Second)

		//resource 1
		if !r2.Lock.TryLock() {
			fmt.Printf("Could not acquire lock for resource %d\n", r2.ID)
			r1.Lock.Unlock()
			continue
		}

		fmt.Printf("Performing operations with resources %d and %d\n", r1.ID, r2.ID)

		time.Sleep(2 * time.Second)

		r1.Lock.Unlock()
		r2.Lock.Unlock()

		time.Sleep(1 * time.Second)
	}
}

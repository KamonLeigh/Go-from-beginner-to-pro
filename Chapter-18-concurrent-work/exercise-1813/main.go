package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
)

func generateRandonNumber(max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))

	if err != nil {
		return 0, err
	}

	return int(n.Int64()), nil
}

func updateCount(countMap *sync.Map, key int) {
	count, _ := countMap.LoadOrStore(key, 0)
	countMap.Store(key, count.(int)+1)
}

func printCounts(countMap *sync.Map) {
	countMap.Range(func(key, value interface{}) bool {
		fmt.Printf("Number %d: Count %d\n", key, value)
		return true
	})
}

func main() {
	var countMap sync.Map
	numGoroutinses := 5
	var wg sync.WaitGroup

	generateAndCount := func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			// Generate random number between 0 and 9
			randomnNumber, err := generateRandonNumber(10)

			if err != nil {
				fmt.Println("Error generating rando, number:", err)
				return
			}

			updateCount(&countMap, randomnNumber)
		}
	}

	for i := 0; i < numGoroutinses; i++ {
		wg.Add(1)
		go generateAndCount()
	}

	wg.Wait()
	printCounts(&countMap)

}

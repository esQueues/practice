package main

import (
	"fmt"
	"github.com/fatih/color"
	"sync"
	"time"
)

func main() {
	numbers := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		numbers[i] = i + 1
	}

	startTime := time.Now()
	for i := range numbers {
		factorial(numbers[i])
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	color.Red("БЕЗ ГОУ РУТИН: %v\n", elapsedTime)

	goroutines := []int{1, 4, 8, 10, 100, 1000}

	for _, numGoroutines := range goroutines {
		wg := sync.WaitGroup{}

		startTime := time.Now()
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				for j := i * 100; j < min((i+1)*100, len(numbers)); j++ {
					factorial(numbers[j])
				}
			}(i)
		}

		wg.Wait()
		endTime := time.Now()
		elapsedTime := endTime.Sub(startTime)
		fmt.Printf("Время с %d горутинам: %v\n", numGoroutines, elapsedTime)
	}
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

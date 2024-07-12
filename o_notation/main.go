package main

import (
	"fmt"
	"sync"
)

func MinEl2(a []int) int {
	// только для первичной проверки
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}

	midIndex := len(a) / 2
	var wg sync.WaitGroup
	wg.Add(2)

	var t1, t2 int

	go func() {
		defer wg.Done()
		t1 = MinEl2(a[:midIndex])
	}()

	go func() {
		defer wg.Done()
		t2 = MinEl2(a[midIndex:])
	}()

	wg.Wait()

	if t1 <= t2 {
		return t1
	}
	return t2
}

func main() {
	a := []int{16, 6, 6, 10, 15, 65, 12}
	fmt.Println(MinEl2(a))
}

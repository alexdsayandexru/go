package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type chantype interface{}

type Value struct {
	i int
	v int
}

func first(in chan chantype) {
	i := 1
	for value := rand.Intn(100); value != 50; value = rand.Intn(100) {
		in <- Value{i, value}
		in <- value
		i++
	}
}

func second(out chan chantype) {
	for values := range out {
		fmt.Println(values)
	}
}

func main() {
	c1 := make(chan chantype)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		first(c1)
		close(c1)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		second(c1)
	}()
	wg.Wait()
}

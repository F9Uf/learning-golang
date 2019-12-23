package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string, 10)
	go findImages(ch)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for filename := range ch {
				resizeImage(filename)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func findImages(ch chan string) {
	for i := 1; i <= 20; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Found file-%d\n", i)
		ch <- fmt.Sprintf("file-%d", i)
	}
	close(ch)
}

func resizeImage(filename string) {
	fmt.Printf("resize image: %s\n", filename)
	time.Sleep(2 * time.Second)
}

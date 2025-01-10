package main

import (
	"fmt"
	"sync"
)

// Menerima data melalui channel dan mencetaknya
func processData(ch chan string, wg *sync.WaitGroup) {
	data := <-ch
	fmt.Println("Data received:", data)
	wg.Done()
}

var (
	state int
	mutex sync.Mutex
)

func incrementState(wg *sync.WaitGroup) {
	mutex.Lock()
	defer mutex.Unlock()

	state++
	fmt.Println("State after increment:", state)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(6)
	// Menggunakan Channels untuk Memindahkan Kepemilikan Data
	// Membuat channel untuk string
	dataChannel := make(chan string)

	// Mengirim data ke channel (memindahkan kepemilikan data)
	go processData(dataChannel, &wg)
	dataChannel <- "Hello, World!" // Data dikirim ke channel dan kepemilikannya berpindah

	// Menggunakan Mutexes untuk Mengelola State
	for i := 0; i < 5; i++ {
		go incrementState(&wg)
	}

	wg.Wait()

	fmt.Println("Done")
}

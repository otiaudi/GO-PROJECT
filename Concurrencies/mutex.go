package main

import (
    "fmt"
    "sync"
)

var count = 0              // Shared whiteboard
var mutex = sync.Mutex{}   // The marker pass

func write(name string, wg *sync.WaitGroup) {
    mutex.Lock()           // Take the marker
    count++
    fmt.Println(name, "wrote. Total =", count)
    mutex.Unlock()         // Give the marker back
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go write("Alice", &wg)
    go write("Bob", &wg)

    wg.Wait()
    fmt.Println("✅ All done!")
}

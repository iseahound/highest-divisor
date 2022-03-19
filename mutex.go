package main

import (
    "fmt"
    "sync"
    "time"
)

type Container struct {
    mu       sync.Mutex
    counters map[int]int
}

func (c *Container) inc(quotient int) {
    // Mutex lock.
    c.mu.Lock()

    // Unlock mutex after function exits.
    defer c.mu.Unlock()

    // increment the shared memory.
    c.counters[quotient]++
}

func main() {

    // Define the range from 1-10000.
    max_value := 10000
    max_divisor := 100

    // Create a frequency map.
    frequency := make(map[int]int)

    // Initialize frequency map values to zero.
    for i := 1; i <= max_value; i++ {
        frequency[i] = 0
    }

    // Create an instance of the container struct.
    c := Container{
        // mutex default value is fine.
        counters: frequency, // Use the frequency map.
    }

    // Create a wait for multiple objects instance.
    var wg sync.WaitGroup
    wg.Add(max_divisor)

    // Create a function that checks the divisors
    doIncrement := func(divisor int, max_value int) {
        for i := 1; i <= max_value; i++ {
            if i % divisor == 0 {
                c.inc(i)
            }
        }
        // Object is complete.
        wg.Done()
    }


    // Start Time
    start := time.Now()

    for divisor := 1; divisor <= max_divisor; divisor++ {
        go doIncrement(divisor, max_value)
    }

    // Allow each thread (goroutine) to finish.
    wg.Wait()
    t := time.Now() // Capture end time.


    // Print time.
    fmt.Println("Time Elapsed:")
    fmt.Println(t.Sub(start))

    // Grab the number of divisors.
    most_divisors := 0
    number_of_divisors := 0
    for i := 1; i <= max_value; i++ {
        if c.counters[i]  > number_of_divisors {
            number_of_divisors = c.counters[i]
            most_divisors = i
        }
    }
    fmt.Println("The number with the most divisors is:")
    fmt.Println(most_divisors)
    fmt.Println("with a frequency of:")
    fmt.Println(number_of_divisors)

}

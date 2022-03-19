package main

import (
    "fmt"
    "time"
)

func worker(done chan int, max_value int, divisor int) {
    for i := 1; i <= max_value; i++ {
        if i % divisor == 0 {
            done <- i
        }
    }
    close(done)
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

    // Create an array of channels.
    channels := make([]chan int, max_divisor)

    // Start time.
    start := time.Now()

    // Launch 10000 goroutines.
    for i, _ := range channels {
        channels[i] = make(chan int, max_value)
        go worker(channels[i], max_value, i+1)
    }

    // Wait for all go routines to finish.
    // Blocking.
    for i, _ := range channels {
        for quotient := range channels[i] {
            frequency[quotient] += 1
        }
    }

    // Print time.
    t := time.Now()
    fmt.Println("Time Elapsed:")
    fmt.Println(t.Sub(start))

    // Grab the number of divisors.
    most_divisors := 0
    number_of_divisors := 0
    for i := 1; i <= max_value; i++ {
        if frequency[i] > number_of_divisors {
            number_of_divisors = frequency[i]
            most_divisors = i
        }
    }
    fmt.Println("The number with the most divisors is:")
    fmt.Println(most_divisors)
    fmt.Println("with a frequency of:")
    fmt.Println(number_of_divisors)









    // Initialize the map with frequencies of zero.
    for i := 1; i <= max_value; i++ {
        frequency[i] = 0
    }

    most_divisors = 0
    number_of_divisors = 0

    start = time.Now()

    for divisor := 1; divisor <= max_divisor; divisor++ {
        for i := 1; i <= max_value; i++ {
            if i % divisor == 0 {
                frequency[i] += 1
                if frequency[i] > number_of_divisors {
                    number_of_divisors = frequency[i]
                    most_divisors = i
                }
            }
        }
    }

    t = time.Now()
    fmt.Println("Time Elapsed:")
    fmt.Println(t.Sub(start))


    fmt.Println("The number with the most divisors is:")
    fmt.Println(most_divisors)
    fmt.Println("with a frequency of:")
    fmt.Println(number_of_divisors)

    // fmt.Println(frequency)
}

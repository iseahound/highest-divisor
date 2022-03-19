# highest-divisor
Computing the number with the highest number of divisors in Go

This compares a mutex vs go channels. The mutex takes about 12 ms, the goroutines + channels 4 ms, and the synchronous control group is about 5 ms. I suspect that the large amount of lock() and unlock() operations on the mutex contributed greatly to its inefficiency.

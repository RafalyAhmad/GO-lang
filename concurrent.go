// goroutine dan channel

func recursiveFunction(n int) int {
    // Base case
    if n == 0 {
        return 1
    }

    // Recursive case
    return n * recursiveFunction(n-1)
}
package main

func concurrentFib(n int) []int {
	ans := make([]int, 0)

	ch := make(chan int)

	go fibonacci(n, ch)

	for i := range ch {
		ans = append(ans, i)
	}
	return ans
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

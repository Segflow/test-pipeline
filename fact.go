package main

func factRecursive(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 1
	}
	return n * factRecursive(n-1)
}

func factIter(n int) int {
	if n < 0 {
		return -1
	}

	a := 1
	for i := 1; i <= n; i++ {
		a = a * i
	}

	return a
}

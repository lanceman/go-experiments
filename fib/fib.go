package fib

func Add(a, b uint64) uint64 {
	return a + b
}

func NextFib(start uint64) uint64 {
	var fib, prev uint64 = 1, 0
	for i := prev; ; {
		if fib > start {
			return fib
		}
		prev = fib
		fib = prev + i
		i = prev
	}
}

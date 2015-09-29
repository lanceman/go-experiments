package fib

func NextFib(start uint64) (fNum, sNum uint64) {
	var prev uint64 = 0
	fNum, sNum = 1,1
	for i := prev; ; {
		if fNum > start {
			return fNum, sNum
		}
		prev = fNum
		fNum = prev + i
		i = prev
		sNum++ 
	}
}

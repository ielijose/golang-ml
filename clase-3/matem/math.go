package matem

func Add(x, y int) int {
	return x + y
}


func Sum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum = sum + i
		}
	}
	return sum	
}
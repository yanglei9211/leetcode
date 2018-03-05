package main

func main() {
	x := 34.00515
	y := -3
	fmt.Println(myPow(x, y))
}

func myPow(x float64, n int) float64 {
	var flag bool
	if n == 0 {
		return 1
	} else if n < 0 {
		flag = true
		n = -n
	} else {
		flag = false
	}

	t := 1.0
	rx := x
	for n > 0 {
		if (n & 1) > 0 {
			t = t * rx
		}
		rx = rx * rx
		n >>= 1
	}
	if flag {
		return 1.0 / t
	} else {
		return t
	}
}

package main

import "fmt"

func area(x1, y1, x2, y2, x3, y3 int) float64 {
	s := (x1*y2 + x2*y3 + x3*y1 - x1*y3 - x2*y1 - x3*y2)
	res := 0.5 * float64(s)
	return res
}

func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	var ans float64 = 0.0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				c := area(points[i][0], points[i][1], points[j][0], points[j][1], points[k][0], points[k][1])
				if c < 0 {
					c = -c
				}
				if ans < c {
					ans = c
				}
			}
		}
	}
	return ans
}

func main() {
	a := [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}
	fmt.Println(largestTriangleArea(a))
}

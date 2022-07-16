package main

import "fmt"

// func maxPoints(points [][]int) int {
// 	line := map[string]int{}
// 	m := 0
// 	for i := 0; i < len(points)-1; i++ {
// 		for j := i + 1; j < len(points); j++ {
// 			x := points[j][0] - points[i][0]
// 			y := points[j][1] - points[i][1]
// 			key := ""

// 			switch {
// 			case x == 0:
// 				key = fmt.Sprintf("_%d", points[j][0])
// 			case y == 0:
// 				key = fmt.Sprintf("0_%d", points[j][1])
// 			default:
// 				slope := float64(y) / float64(x)
// 				interpret := float64(points[i][1]) - slope*float64(points[i][0])
// 				key = fmt.Sprintf("%f_%f", slope, interpret)
// 			}

// 			if value, ok := line[key]; ok {
// 				line[key] = value + 1
// 			} else {
// 				line[key] = 1
// 			}

// 			if m < line[key] {
// 				m = line[key]
// 			}
// 		}
// 	}

// 	n := (1.0 + math.Sqrt(1.0+8.0*float64(m))) / 2.0
// 	return int(n)
// }

// func maxPoints(points [][]int) int {
// 	gcd := func(a, b int) int {
// 		if b < 0 {
// 			b = -b
// 		}
// 		if a < 0 {
// 			a = -a
// 		}
// 		if b > a {
// 			a, b = b, a
// 		}
// 		for b > 0 {
// 			r := a % b
// 			a, b = b, r
// 		}
// 		return a
// 	}

// 	if len(points) == 1 {
// 		return 1
// 	}
// 	type line struct {
// 		a int
// 		b int
// 		c int
// 	}
// 	type unit struct{}
// 	t := make(map[line]map[int]unit, len(points)*(len(points)-1)/2)
// 	var result int
// 	for i, p1 := range points {
// 		x1, y1 := p1[0], p1[1]
// 		for j, p2 := range points[i+1:] {
// 			x2, y2 := p2[0], p2[1]

// 			a := y2 - y1
// 			b := x1 - x2
// 			c := x1*(y1-y2) + y1*(x2-x1)
// 			d := gcd(a, gcd(b, c))
// 			a /= d
// 			b /= d
// 			c /= d

// 			line := line{a, b, c}

// 			tt, ok := t[line]
// 			if !ok {
// 				tt = make(map[int]unit)
// 				t[line] = tt
// 			}
// 			tt[i] = unit{}
// 			tt[i+1+j] = unit{}
// 			if count := len(tt); count > result {
// 				result = count
// 			}
// 		}
// 	}
// 	return result
// }

// func maxPoints(points [][]int) int {
// 	p := points
// 	l := len(p)
// 	v := make([]map[int]int, l)
// 	v0, v8 := make([]int, l), make([]int, l)
// 	ans := 1
// 	for i := 0; i < l; i++ {
// 		v[i] = map[int]int{}
// 	}
// 	for i := 0; i < l && l-i > ans; i++ {
// 		for j := i + 1; j < l; j++ {
// 			vx, vy := float64(p[j][0]-p[i][0]), float64(p[j][1]-p[i][1])
// 			if vx == 0 {
// 				if v8[j] > 0 {
// 					continue
// 				} else {
// 					v8[j] = 1
// 					if v8[i] == 0 {
// 						v8[i] = 2
// 					} else {
// 						v8[i]++
// 					}
// 				}
// 				if ans < v8[i] {
// 					ans = v8[i]
// 				}
// 				continue
// 			}
// 			if vy == 0 {
// 				if v0[j] > 0 {
// 					continue
// 				} else {
// 					v0[j] = 1
// 					if v0[i] == 0 {
// 						v0[i] = 2
// 					} else {
// 						v0[i]++
// 					}
// 				}
// 				if ans < v0[i] {
// 					ans = v0[i]
// 				}
// 				continue
// 			}
// 			vt := int((vy / vx) * 1e4)
// 			if _, ok := v[j][vt]; ok {
// 				continue
// 			}
// 			if _, ok := v[i][vt]; ok {
// 				v[i][vt]++
// 			} else {
// 				v[i][vt], v[j][vt] = 2, 1
// 			}
// 			if ans < v[i][vt] {
// 				ans = v[i][vt]
// 			}
// 		}
// 	}
// 	return ans
// }

type pair struct {
	dy int
	dx int
}

func maxPoints(points [][]int) int {
	ret := 0
	n := len(points)
	for i := 0; i < n; i += 1 {
		m := make(map[pair]int)
		dup := 1
		curMax := 0
		for j := i + 1; j < n; j += 1 {
			dy := points[j][1] - points[i][1]
			dx := points[j][0] - points[i][0]
			if dx == 0 && dy == 0 {
				dup += 1
			} else {
				key := getSlope(points, i, j)
				if v, ok := m[key]; ok {
					m[key] = v + 1
				} else {
					m[key] = 1
				}
				curMax = max(curMax, m[key])
			}
		}
		ret = max(ret, curMax+dup)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getSlope(points [][]int, i, j int) pair {
	dy := points[j][1] - points[i][1]
	dx := points[j][0] - points[i][0]
	if dy == 0 {
		return pair{0, points[i][1]}
	}
	if dx == 0 {
		return pair{points[i][0], 0}
	}
	d := gcd(dy, dx)
	return pair{dy / d, dx / d}
}

func gcd(y, x int) int {
	if x == 0 {
		return y
	}
	return gcd(x, y%x)
}

type data struct {
	input    [][]int
	expected int
}

func main() {
	sols := []data{
		{[][]int{{1, 1}, {2, 2}, {3, 3}}, 3},
		{[][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}, 4},
		{[][]int{{2, 3}, {3, 3}, {-5, 3}}, 3},
		{[][]int{{7, 3}, {19, 19}, {-16, 3}, {13, 17}, {-18, 1}, {-18, -17}, {13, -3}, {3, 7},
			{-11, 12}, {7, 19}, {19, -12}, {20, -18}, {-16, -15}, {-10, -15}, {-16, -18}, {-14, -1},
			{18, 10}, {-13, 8}, {7, -5}, {-4, -9}, {-11, 2}, {-9, -9}, {-5, -16}, {10, 14}, {-3, 4},
			{1, -20}, {2, 16}, {0, 14}, {-14, 5}, {15, -11}, {3, 11}, {11, -10}, {-1, -7}, {16, 7},
			{1, -11}, {-8, -3}, {1, -6}, {19, 7}, {3, 6}, {-1, -2}, {7, -3}, {-6, -8}, {7, 1},
			{-15, 12}, {-17, 9}, {19, -9}, {1, 0}, {9, -10}, {6, 20}, {-12, -4}, {-16, -17},
			{14, 3}, {0, -1}, {-18, 9}, {-15, 15}, {-3, -15}, {-5, 20}, {15, -14}, {9, -17},
			{10, -14}, {-7, -11}, {14, 9}, {1, -1}, {15, 12}, {-5, -1}, {-17, -5}, {15, -2},
			{-12, 11}, {19, -18}, {8, 7}, {-5, -3}, {-17, -1}, {-18, 13}, {15, -3}, {4, 18},
			{-14, -15}, {15, 8}, {-18, -12}, {-15, 19}, {-9, 16}, {-9, 14}, {-12, -14}, {-2, -20},
			{-3, -13}, {10, -7}, {-2, -10}, {9, 10}, {-1, 7}, {-17, -6}, {-15, 20}, {5, -17},
			{6, -6}, {-11, -8}}, 6},
	}
	fmt.Println(sols)

	for i, sol := range sols {
		if output := maxPoints(sol.input); sol.expected != output {
			fmt.Println(i, "output: ", output, "expected: ", sol.expected)
		}
	}
}
